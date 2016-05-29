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
	"bufio"
	"fmt"
	"strconv"
	"time"
)

var L = utils.L

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

// // 服务器信息
// type ServerInfo struct {
// 	ip          string
// 	port        int32
// 	currentLoad int32
// 	serverType  protocol.ClusterServerType
// }

// type ServerList struct {
// 	serverType protocol.ClusterServerType
// 	data       map[string]ServerInfo
// }

type AppConfig struct {
	serverPort  int32
	clusterIp   string
	clusterPort int32
}

var appConfig AppConfig

// 服务器列表
var Servers = make(map[string]protocol.ServerInfo)

// var selfInfo ServerInfo

// 初始化配置
func initConf() {
	appConfig = AppConfig{}
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	utils.CheckError(err)
	var v int
	v, err = appConf.Int("server::port")
	utils.CheckError(err)
	appConfig.serverPort = int32(v)
	appConfig.clusterIp = appConf.String("cluster::ip")
	v, err = appConf.Int("cluster::port")
	utils.CheckError(err)
	appConfig.clusterPort = int32(v)
}

func Run() {
	L.Info("Login 服务启动")

	// 初始化配置
	initConf()

	// 连接到 cluster服务器
	go connectCluster()

	// go count()
	// go schedule()
	addr := ":" + strconv.Itoa(int(appConfig.serverPort))
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	utils.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)
	L.Info("服务监听端口:%d", appConfig.serverPort)
	for {
		L.Info("当前负载：%d", len(coroutinePool))
		conn, err := listener.Accept()
		if err != nil {
			L.Error("请求失败")
			continue
		}
		if err != nil {
			L.Error("%d", len(coroutinePool))
		}
		// 控制协程数量
		coroutinePool <- 1
		go func(con net.Conn) {
			handleClient(con)
			<-coroutinePool
		}(conn)

	}
}

// 连接到 Cluster 服务器上
func connectCluster() {
	fmt.Println("已连接到Cluster服务器", appConfig.clusterPort)
	L.Debug("已连接到Cluster服务器", appConfig.clusterPort)
	service := appConfig.clusterIp + ":" + strconv.Itoa(int(appConfig.clusterPort))
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		utils.CheckError(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		utils.CheckError(err)
	}
	defer conn.Close()
	L.Debug("已连接到Cluster服务器")
	c := make(chan bool)
	go regServer(conn)
	go scheduleReg(conn)
	<-c
}

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

func scheduleReg(conn *net.TCPConn) {

}

// func onMessageRecived(conn *net.TCPConn) {
//     reader := bufio.NewReader(conn)
//     for {
//         msg, err := reader.ReadString('\n')
//         fmt.Println(msg)
//         if err != nil {
//             quitSemaphore <- true
//             break
//         }
//         time.Sleep(time.Second)
//         b := []byte(msg)
//         conn.Write(b)
//     }
// }

// 注册自身
func regServer(conn *net.TCPConn) {
	L.Debug("注册自身")
	_, err := conn.Write(getServerResponse())
	if err != nil {
		utils.CheckError(err)
	}
	reader := bufio.NewReader(conn)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}
	}
	// for {
	// 	reader := bufio.NewReader(conn)
	// 	buf := make([]byte, 2048)
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		utils.CheckError(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	rdata := &protocol.ClusterResponse{}
	// 	err = proto.Unmarshal(buf[:n], rdata)
	// 	if err != nil {
	// 		break
	// 		// L.Debug("连接到Cluster服务器")
	// 		// utils.CheckError(err)
	// 	}
	// }
}

func heartbeet(conn net.Conn) {

}

// 定时发送带负载信息的心跳包
func scheduleHeartbeat(conn net.Conn) {
	t := time.NewTimer(internal)
	for {
		<-t.C
		heartbeet(conn)
		t.Reset(internal)
		// break
	}
}

// 当前服务器信息
func getServerResponse() []byte {
	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_REG_SERVER,
		Data: []*protocol.ServerInfo{
			{
				Type:        protocol.ServerType_LOGIN,
				Ip:          appConfig.clusterIp,
				Port:        appConfig.clusterPort,
				CurrentLoad: int32(len(coroutinePool)),
			},
		},
	}
	data, err := proto.Marshal(lres)
	utils.CheckError(err)
	return data
}

// 定时检查所有服务器最后心跳时间是否超过1分钟，超过的话则为未存活，因为所有服务器间隔30秒要向服务器发送心跳包
func schedule() {
	// t := time.NewTimer(internal)
	// for {
	// 	<-t.C
	// 	// regServer()
	// 	t.Reset(internal)
	// 	// break
	// }
}
