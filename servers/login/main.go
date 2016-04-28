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

	lres := &protocol.LoginResponse{
		status: Status_SUCCESS,
	}

	data, err := proto.Marshal(lres)
	checkError(err)

	// var res LoginResponse
	// res.status = Status_SUCCESS
	L.Debug("%s", res.status)

	service := ":9001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
	L.Info("Serving at localhost:9001")
	// log.Fatal(http.ListenAndServe(":9001", nil))
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			L.Error(err.Error())
			break
		}
		L.Debug(string(request))
		if read_len == 0 {
			break
		} else {
			conn.Write([]byte(time.Now().String()))
		}
		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		L.Critical("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
