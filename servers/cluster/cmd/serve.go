package cmd

import (
	"github.com/astaxie/beego/config"
	"github.com/codegangsta/cli"
	"github.com/golang/protobuf/proto"
	// "github.com/google/uuid"
	"github.com/xuhuan/keepin/protocol"
	"github.com/xuhuan/keepin/utils"
	// "log"
	"net"
	// "os"
	// "runtime"
	"strconv"
	"time"
)

var L = utils.L

var aliveTimeout = time.Second * 60
var internal = time.Second * 30
var dateFormat = "2006-01-02 15:04:05"

var CmdServe = cli.Command{
	Name:        "serve",
	Usage:       "start server cluster ",
	Description: `server for servers`,
	Action:      runServe,
	Flags: []cli.Flag{
		utils.StringFlag("config, c", "conf/app.conf", "configuration file path"),
	},
}

// 服务器信息
type ServerInfo struct {
	alive             bool
	ip                string
	port              int32
	currentLoad       int32
	serverType        protocol.ClusterServerType
	lastHeartbeatTime string
}

type ServerList struct {
	serverType protocol.ClusterServerType
	data       map[string]ServerInfo
}

// 服务器列表
var Servers = make(map[string]ServerList)

func runServe(ctx *cli.Context) {
	L.Info("Cluster 服务启动")
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	utils.CheckError(err)

	addr := "localhost:" + appConf.String("server::port")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)
	L.Info("服务监听端口:%s", appConf.String("server::port"))
	for {
		conn, err := listener.Accept()
		if err != nil {
			L.Error("请求失败")
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	L.Debug("收到请求")
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 1024)
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			L.Error(err.Error())
			break
		}
		L.Debug("收到消息,长度：%d", read_len)

		data := request[:read_len]
		encode := &protocol.ClusterRequest{}
		err = proto.Unmarshal(data, encode)
		utils.CheckError(err)
		if read_len == 0 {
			break
		}

		switch encode.Act {
		case protocol.ClusterActionType_GET_SERVERS:
			conn.Write(getServer(encode.ServerType))
		case protocol.ClusterActionType_REG_SERVER:
			conn.Write(regServer(encode))
		case protocol.ClusterActionType_HEARTBEAT:
			conn.Write(heartbeat(encode))
		default:
			conn.Write([]byte(time.Now().String()))
		}

		request = make([]byte, 1024)
	}
}

// 更新心跳时间
func heartbeat(server *protocol.ClusterRequest) []byte {
	if server.GetData() != nil {
		serverInfo := server.GetData()[0]
		_, ok := Servers[serverInfo.Type.String()]
		if ok {
			Servers[serverInfo.Type.String()].data[utils.Md5(serverInfo.Ip+":"+strconv.Itoa(int(serverInfo.Port)))] = ServerInfo{
				alive:             true,
				ip:                serverInfo.Ip,
				port:              serverInfo.Port,
				serverType:        serverInfo.Type,
				currentLoad:       serverInfo.CurrentLoad,
				lastHeartbeatTime: time.Now().Format(dateFormat),
			}
		}
	}

	lres := &protocol.ClusterResponse{
		Status: protocol.Status_SUCCESS,
	}
	wdata, err := proto.Marshal(lres)
	utils.CheckError(err)
	return wdata
}

// 输出服务器状态
func printStatus() {
	for k, v := range Servers {
		L.Debug("服务器类型：%s，主机数量：%d", k, len(v.data))
		isLive := 0
		for _, server := range v.data {
			L.Debug("服务器状态：\nIP：%s\n端口：%d\n存活状态:%t\n负载：%d", server.ip, server.port, server.alive, server.currentLoad)
			if server.alive {
				isLive++
			}
		}
		L.Debug("服务器存活数量：%d,非活跃数量：%d", isLive, len(v.data)-isLive)
	}
}

// 获取多个指定类型的负载最小服务器信息
func getServer(serverTypes []protocol.ClusterServerType) []byte {
	data := []*protocol.ClusterServerInfo{}
	for _, serverType := range serverTypes {
		var _server ServerInfo
		i := 0
		for _, server := range Servers[serverType.String()].data {
			if i == 0 {
				_server = server
			} else {
				if server.currentLoad < _server.currentLoad {
					_server = server
				}
			}
			i++
		}
		data = append(data, &protocol.ClusterServerInfo{
			Type:              _server.serverType,
			Ip:                _server.ip,
			Port:              _server.port,
			CurrentLoad:       _server.currentLoad,
			LastHeartbeatTime: _server.lastHeartbeatTime,
		})
	}
	lres := &protocol.ClusterResponse{
		Status: protocol.Status_SUCCESS,
		Data:   data,
	}
	wdata, err := proto.Marshal(lres)
	utils.CheckError(err)
	return wdata
}

// 注册服务器
func regServer(server *protocol.ClusterRequest) []byte {
	if server.GetData() != nil {
		serverInfo := server.GetData()[0]
		_, ok := Servers[serverInfo.Type.String()]
		if !ok {
			Servers[serverInfo.Type.String()] = ServerList{
				serverType: serverInfo.Type,
				data:       make(map[string]ServerInfo),
			}
		}
		Servers[serverInfo.Type.String()].data[utils.Md5(serverInfo.Ip+":"+strconv.Itoa(int(serverInfo.Port)))] = ServerInfo{
			alive:             true,
			ip:                serverInfo.Ip,
			port:              serverInfo.Port,
			serverType:        serverInfo.Type,
			currentLoad:       serverInfo.CurrentLoad,
			lastHeartbeatTime: time.Now().Format(dateFormat),
		}
	}
	// printStatus()

	lres := &protocol.ClusterResponse{
		Status: protocol.Status_SUCCESS,
	}
	wdata, err := proto.Marshal(lres)
	utils.CheckError(err)
	return wdata
}

// 定时检查所有服务器最后心跳时间是否超过1分钟，超过的话则为未存活，因为所有服务器间隔30秒要向服务器发送心跳包
func schedule() {
	t := time.NewTimer(internal)
	for {
		<-t.C
		checkAlive()
		t.Reset(internal)
		// break
	}
}

// 检测是否存活
func checkAlive() {
	for k, v := range Servers {
		L.Debug("服务器类型：%s，主机数量：%d，检测开始...", k, len(v.data))
		isLive := 0
		for id, server := range v.data {
			t, err := time.Parse(dateFormat, server.lastHeartbeatTime)
			utils.CheckError(err)
			t = t.Add(aliveTimeout)
			now, _ := time.Parse(dateFormat, time.Now().Format(dateFormat))
			if t.Before(now) {
				server.alive = false
			}
			if server.alive {
				isLive++
			}
			v.data[id] = server
		}
		Servers[k] = v
		L.Debug("服务器存活数量：%d,非活跃数量：%d", isLive, len(v.data)-isLive)
	}
}

func Run() {
	L.Info("Cluster 服务启动")
	L.Info(time.Now().Format(dateFormat))
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	utils.CheckError(err)

	go schedule()
	addr := ":" + appConf.String("server::port")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)
	L.Info("服务监听端口:%s", appConf.String("server::port"))
	for {
		conn, err := listener.Accept()
		if err != nil {
			L.Error("请求失败")
			continue
		}
		go handleClient(conn)
	}
}
