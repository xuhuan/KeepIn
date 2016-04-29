package main

import (
	// "github.com/codegangsta/cli"
	// "github.com/xuhuan/keepin/servers/login/client"
	// "github.com/xuhuan/keepin/servers/login/cmd"
	"github.com/xuhuan/keepin/utils"
	// "os"
	// "github.com/googollee/go-socket.io"
	// "log"
	"github.com/golang/protobuf/proto"
	"github.com/xuhuan/keepin/protocol"
	"net"
	"os"
	"runtime"
	// "strconv"
	"log"
	"time"
)

const APP_VER = "0.0.1"

var L = utils.L

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// app := cli.NewApp()
	// app.Name = "KeepIn IM"
	// app.Usage = "KeepIn Login Service"
	// app.Version = APP_VER
	// app.Commands = []cli.Command{
	// 	cmd.CmdWeb,
	// }
	// app.Flags = append(app.Flags, []cli.Flag{}...)
	// app.Run(os.Args)
	L.Info("服务启动")
	// client.Init()
	// cmd.Run()
	// server, err := socketio.NewServer(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// server.On("connection", func(so socketio.Socket) {
	// 	L.Info("连接成功")
	// 	so.Join("room")
	// 	so.On("chat message", func(msg string) {
	// 		L.Info("emit:", so.Emit("chat message", msg))
	// 		so.BroadcastTo("chat", "chat message", msg)
	// 	})
	// 	so.On("disconnection", func() {
	// 		L.Info("连接断开")
	// 	})
	// })
	// server.On("error", func(so socketio.Socket, err error) {
	// 	L.Error("error:", err)
	// })

	// http.Handle("/", server)

	// lres := &protocol.LoginResponse{
	// 	Status:  protocol.Status_SUCCESS,
	// 	Message: "成功",
	// 	Data: &protocol.LoginData{
	// 		ServerTime: 1,
	// 		UserInfo: &protocol.Info{
	// 			Uid:       1,
	// 			Gender:    1,
	// 			NickName:  "昵称",
	// 			AvatarUrl: "http://www.tzrl.com",
	// 		},
	// 	},
	// }

	// data, err := proto.Marshal(lres)
	// checkError(err)
	// L.Info(strconv.Itoa(len(data)))

	// encode := &protocol.LoginResponse{}
	// err = proto.Unmarshal(data, encode)
	// checkError(err)
	// L.Debug("%s %d", encode.Message, encode.Status)
	// L.Debug("%s %d", encode.GetData().GetUserInfo().NickName, encode.GetData().ServerTime)

	// var res LoginResponse
	// res.status = Status_SUCCESS
	// L.Debug("%s", res.status)

	service := ":9002"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	L.Info("服务监听端口:9002")
	for {
		conn, err := listener.Accept()
		if err != nil {
			L.Error("请求失败")
			continue
		}
		go handleClient(conn)
	}
	// log.Fatal(http.ListenAndServe(":9001", nil))
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
		encode := &protocol.LoginResponse{}
		err = proto.Unmarshal(data, encode)
		checkError(err)
		L.Debug("%s %d", encode.Message, encode.Status)
		L.Debug("%s %d", encode.GetData().GetUserInfo().NickName, encode.GetData().ServerTime)
		// L.Debug(string(request))
		if read_len == 0 {
			break
		} else {
			conn.Write([]byte(time.Now().String()))
		}
		request = make([]byte, 1024)
	}
}

func checkError(err error) {
	if err != nil {
		L.Critical("Fatal error: %s", err.Error())
		log.Fatal(err)
		os.Exit(1)
	}
}
