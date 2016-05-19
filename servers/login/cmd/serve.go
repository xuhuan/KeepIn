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
	"runtime"
	// "strconv"
	"fmt"
	"time"
)

var L = utils.L

// 当前负载
var currentLoad = 0

// 协程池大小
var coroutinePool = make(chan int, 100000)

// 定时注册自身间隔
var internal = time.Second * 30

// 时间格式
var dateFormat = "2006-01-02 15:04:05"

// 命令
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
	ip          string
	port        int32
	currentLoad int32
	serverType  protocol.ClusterServerType
}

type ServerList struct {
	serverType protocol.ClusterServerType
	data       map[string]ServerInfo
}

// [server]
// port = 9000

// [cluster]
// ip   =localhost
// port =9200

// type Config struct {
// 	server
// }

// 服务器列表
var Servers = make(map[string]ServerList)

// var selfInfo ServerInfo

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
		// 控制协程数量
		coroutinePool <- 1
		go func(con net.Conn) {
			handleClient(con)
			<-coroutinePool
		}(conn)
	}
}

// func getSelf() {
// 	selfInfo
// }

func handleClient(conn net.Conn) {
	defer conn.Close()
	L.Debug("收到请求")
	// 设置超时时间
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// 设置buff
	buff := make([]byte, 1024)
	for {
		readLen, err := conn.Read(buff)
		if err != nil {
			L.Error(err.Error())
			break
		}
		L.Debug("收到消息,长度：%d", readLen)

		data := buff[:readLen]
		encode := &protocol.ClusterRequest{}
		err = proto.Unmarshal(data, encode)
		utils.CheckError(err)
		if readLen == 0 {
			break
		}

		switch encode.Act {
		// case protocol.LoginActionType_LOGIN:
		// 	conn.Write(login(encode))
		// 	break
		// case protocol.LoginActionType_LOGIN_OUT:
		// 	conn.Write(loginOut(encode))
		// 	break
		default:
			eres := &protocol.LoginResponse{
				Status: protocol.Status_FAIL,
			}
			wdata, err := proto.Marshal(eres)
			utils.CheckError(err)
			conn.Write(wdata)
			break
		}

		// 清空
		buff = make([]byte, 1024)
	}
}

func login(req *protocol.ClusterRequest) []byte {
	d := make([]byte, 1024)
	return d
}

func loginOut(req *protocol.ClusterRequest) []byte {
	d := make([]byte, 1024)
	return d
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
			Type:        _server.serverType,
			Ip:          _server.ip,
			Port:        _server.port,
			CurrentLoad: _server.currentLoad,
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
func regServer() []byte {
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
	utils.CheckError(err)
	return data
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
	fmt.Println("Login 服务启动")
	go func() {
		currentLoad++
		coroutinePool <- 1
		fmt.Println(len(coroutinePool), cap(coroutinePool), runtime.NumGoroutine(), currentLoad)
	}()
	currentLoad--
	fmt.Println(currentLoad)
	<-coroutinePool

	L.Info("Login 服务启动")
	L.Info("%d %d", len(coroutinePool), cap(coroutinePool))

	// L.Info(time.Now().Format(dateFormat))
	// appConf, err := config.NewConfig("ini", "conf/app.conf")
	// utils.CheckError(err)

	// go count()
	// // go schedule()
	// addr := ":" + appConf.String("server::port")
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	// utils.CheckError(err)
	// listener, err := net.ListenTCP("tcp", tcpAddr)
	// utils.CheckError(err)
	// L.Info("服务监听端口:%s", appConf.String("server::port"))
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		L.Error(currentLoad// }
}
