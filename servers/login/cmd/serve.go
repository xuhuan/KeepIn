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

// 当前负载
var currentLoad = 0

var maxLoad = 10

var cl = make(chan int, maxLoad)

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
	alive       bool
	ip          string
	port        int32
	currentLoad int32
	serverType  protocol.ClusterServerType
}

type ServerList struct {
	serverType protocol.ClusterServerType
	data       map[string]ServerInfo
}

// 服务器列表
var Servers = make(map[string]ServerList)

func runServe(ctx *cli.Context) {
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

func reader(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// 负载值加1
	cl <- 1
for  {
	select{
		case
	}
}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	L.Debug("收到请求")
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
		case protocol.LoginActionType_LOGIN:
			conn.Write(getServer(encode.ServerType))
			break
		case protocol.LoginActionType_LOGIN_OUT:
			conn.Write(regServer(encode))
			break
		default:
			eres := &protocol.LoginResponse{
				Status: protocol.Status_FAIL,
			}
			wdata, err := proto.Marshal(eres)
			utils.CheckError(err)
			conn.Write(wdata)
			break
		}

		request = make([]byte, 1024)
	}
	// 负载值减1
	cl <- -1
}

// 统计当前负载
func count() {
	for {
		currentLoad += <-cl
		L.Debug("当前负载：%d", currentLoad)
	}
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

	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_REG_SERVER,
		Data: []*protocol.ClusterServerInfo{
			{
				Type:        protocol.ClusterServerType_LOGIN,
				Ip:          "188.66.66.133",
				Port:        8888,
				CurrentLoad: 66,
			},
		},
	}
	data, err := proto.Marshal(lres)

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
		regServer()
		t.Reset(internal)
		// break
	}
}

func Run() {
	L.Info("Login 服务启动")
	L.Info(time.Now().Format(dateFormat))
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	utils.CheckError(err)

	go count()
	// go schedule()
	addr := ":" + appConf.String("server::port")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	utils.CheckError(err)
	L.Info("服务监听端口:%s", appConf.String("server::port"))
	for {
		conn, err := listener.Accept()
		if err != nil {
			L.Error("请求失败")
			continue
		}
		go handleClient(conn)
		currentLoad += <-cl
		L.Debug("负载：%d", currentLoad)
	}
}
