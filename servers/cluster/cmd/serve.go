package cmd

import (
	"github.com/astaxie/beego/config"
	"github.com/codegangsta/cli"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
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

var CmdServe = cli.Command{
	Name:        "serve",
	Usage:       "start server cluster ",
	Description: `server for servers`,
	Action:      runServe,
	Flags: []cli.Flag{
		utils.StringFlag("config, c", "conf/app.conf", "configuration file path"),
	},
}

type ServerInfo struct {
	uuid        string
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

// var Servers = []ServerInfo{}
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
		L.Debug("%s", encode.Act)
		serverInfo := encode.GetData()
		if serverInfo != nil {
			L.Debug("%s", serverInfo[0].Ip)
		}
		if read_len == 0 {
			break
		} else {
			switch encode.Act {
			case protocol.ClusterActionType_GET_SERVERS:
				// lres := &protocol.ClusterResponse{
				// 	Status: protocol.Status_SUCCESS,
				// 	Data: []*protocol.ClusterServerInfo{
				// 		{
				// 			Type: protocol.ClusterServerType_LOGIN,
				// 			Ip:   "188.66.66.33",
				// 			Port: 8888,
				// 		},
				// 	},
				// }
				// wdata, err := proto.Marshal(lres)
				// utils.CheckError(err)
				// conn.Write(wdata)
				conn.Write(getServer(encode.ServerType))
			case protocol.ClusterActionType_REG_SERVER:
				regServer(encode)
				conn.Write([]byte(time.Now().String()))
			default:
				conn.Write([]byte(time.Now().String()))

			}
		}
		request = make([]byte, 1024)
	}
}

func printStatus() {
	for k, v := range Servers {
		L.Info("服务器类型：%s，主机数量：%d", k, len(v.data))
		isLive := 0
		for _, server := range v.data {
			L.Info("服务器状态：\nIP：%s\n端口：%d\n存活状态:%t\n负载：%d", server.ip, server.port, server.alive, server.currentLoad)
			if server.alive {
				isLive++
			}
		}
		L.Info("服务器存活数量：%d,非活跃数量：%d", isLive, len(v.data)-isLive)
	}
}

func getServer(serverType protocol.ClusterServerType) []byte {
	L.Debug("%s", serverType.String())
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
	lres := &protocol.ClusterResponse{
		Status: protocol.Status_SUCCESS,
		Data: []*protocol.ClusterServerInfo{
			{
				Type:        _server.serverType,
				Ip:          _server.ip,
				Port:        _server.port,
				CurrentLoad: _server.currentLoad,
				Uuid:        _server.uuid,
			},
		},
	}
	wdata, err := proto.Marshal(lres)
	utils.CheckError(err)
	return wdata
}

func regServer(server *protocol.ClusterRequest) {
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
			uuid:        uuid.New().String(),
			alive:       true,
			ip:          serverInfo.Ip,
			port:        serverInfo.Port,
			serverType:  serverInfo.Type,
			currentLoad: serverInfo.CurrentLoad,
		}

		// if !existServer(serverInfo) {
		// 	L.Debug(serverInfo.Type.String())
		// 	L.Debug(uuid.New().String())
		// 	L.Debug(strconv.Itoa(len(Servers)))
		// 	// L.Debug(strconv.Itoa(len(Servers[serverInfo.Type.String()].data)))
		// 	Servers[serverInfo.Type.String()].data[utils.Md5(serverInfo.Ip+":"+strconv.Itoa(int(serverInfo.Port)))] = ServerInfo{
		// 		uuid:       uuid.New().String(),
		// 		alive:      true,
		// 		ip:         serverInfo.Ip,
		// 		port:       serverInfo.Port,
		// 		serverType: serverInfo.Type,
		// 	}
		// 	// _servers := Servers[serverInfo.Type.String()]
		// 	// _servers.data = append(_servers.data, ServerInfo{
		// 	// 	uuid:       uuid.New().String(),
		// 	// 	alive:      true,
		// 	// 	ip:         serverInfo.Ip,
		// 	// 	port:       serverInfo.Port,
		// 	// 	serverType: serverInfo.Type,
		// 	// })
		// 	// Servers[serverInfo.Type.String()] = _servers
		// 	L.Debug(strconv.Itoa(len(Servers[serverInfo.Type.String()].data)))
		// }
	}
	printStatus()
	// L.Debug(strconv.Itoa(len(Servers)))
	// L.Debug(Servers[0].uuid)
	// L.Debug(strconv.Itoa(int(Servers[0].port)))
}

// // 验证时候已经存在
// func existServer(server *protocol.ClusterServerInfo) bool {
// 	L.Debug(utils.Md5(server.Ip + ":" + strconv.Itoa(int(server.Port))))
// 	if len(Servers[server.Type.String()].data) == 0 {
// 		Servers[server.Type.String()] = ServerList{
// 			serverType: server.Type,
// 			data:       make(map[string]ServerInfo),
// 		}
// 		return false
// 	}
// 	d := Servers[server.Type.String()].data
// 	for i := 0; i < len(d); i++ {
// 		L.Debug(utils.Md5(server.Ip + ":" + strconv.Itoa(int(server.Port))))
// 		_, ok := d[utils.Md5(server.Ip+":"+strconv.Itoa(int(server.Port)))]
// 		return ok
// 		// if ok {
// 		// 	return true
// 		// }
// 	}
// 	return false
// }

func Run() {
	L.Info("Cluster 服务启动")
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	utils.CheckError(err)

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
