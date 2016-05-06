package cmd

import (
	"github.com/astaxie/beego/config"
	"github.com/codegangsta/cli"
	"github.com/golang/protobuf/proto"
	"github.com/xuhuan/keepin/protocol"
	"github.com/xuhuan/keepin/utils"
	// "log"
	"net"
	// "os"
	// "runtime"
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
				lres := &protocol.ClusterResponse{
					Status: protocol.Status_SUCCESS,
					Data: []*protocol.ClusterServerInfo{
						&protocol.ClusterServerInfo{
							Type: protocol.ClusterServerType_LOGIN,
							Ip:   "188.66.66.33",
							Port: 8888,
						},
					},
				}
				wdata, err := proto.Marshal(lres)
				utils.CheckError(err)
				conn.Write(wdata)
			default:
				conn.Write([]byte(time.Now().String()))

			}
		}
		request = make([]byte, 1024)
	}
}

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
