package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/xuhuan/keepin/protocol"
	"github.com/xuhuan/keepin/utils"
	// "io/ioutil"
	"log"
	"net"
	"os"
)

var L = utils.L

func main() {
	service := ":9002"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	L.Debug("连接服务器")

	lres := &protocol.LoginResponse{
		Status:  protocol.Status_SUCCESS,
		Message: "成功",
		Data: &protocol.LoginData{
			ServerTime: 1,
			UserInfo: &protocol.Info{
				Uid:       1,
				Gender:    1,
				NickName:  "昵称",
				AvatarUrl: "http://www.tzrl.com",
			},
		},
	}
	data, err := proto.Marshal(lres)
	checkError(err)

	_, err = conn.Write(data)
	// _, err = conn.Write([]byte("test"))
	checkError(err)
	L.Debug("发送数据")
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		checkError(err)
		if n == 0 {
			break
		}
		L.Debug("接收数据")
		L.Debug(string(buf[:n]))
	}
	// result, err := ioutil.ReadAll(conn)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		L.Critical("Fatal error: %s", err.Error())
		log.Fatal(err)
		os.Exit(1)
	}
}
