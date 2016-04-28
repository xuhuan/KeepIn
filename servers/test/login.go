package main

import (
	"github.com/xuhuan/keepin/protocol/Login"
	"github.com/xuhuan/keepin/utils"
	// "io/ioutil"
	"net"
	"os"
)

var L = utils.L

func main() {
	service := "192.168.0.189:9001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	L.Debug("连接服务器")
	var res LoginResponse
	res.status = Status_SUCCESS
	L.Debug("%s", res.status)
	_, err = conn.Write([]byte("test"))
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
		os.Exit(1)
	}
}
