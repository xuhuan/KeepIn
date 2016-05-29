/*
* @Author: xuhuan
* @Date:   2016-05-06 23:41:49
* @Last Modified by:   xuhuan
* @Last Modified time: 2016-05-30 00:32:49
 */

package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/xuhuan/keepin/protocol"
	"net"
	"strconv"
	"testing"
)

func Test_RequestGetServers(t *testing.T) {
	service := "localhost:9200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("连接服务器")

	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_GET_SERVERS,
	}
	data, err := proto.Marshal(lres)
	t.Log(strconv.Itoa(len(data)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		rdata := &protocol.ClusterResponse{}
		err = proto.Unmarshal(buf[:n], rdata)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("接收数据")
		t.Log(rdata.Status)
		if len(rdata.GetData()) > 0 {
			t.Log(rdata.GetData()[0].Ip)
		}
		t.Log("测试成功")
		break
	}
}

func Test_RequestRegistServer(t *testing.T) {
	service := "localhost:9200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("连接服务器")

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
	t.Log(strconv.Itoa(len(data)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		rdata := &protocol.ClusterResponse{}
		err = proto.Unmarshal(buf[:n], rdata)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("接收数据")
		t.Log(rdata.Status)
		t.Log("测试成功")
		break
	}
}

func Test_RequestRegistLoginServer(t *testing.T) {
	service := "localhost:9200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("连接服务器")

	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_REG_SERVER,
		Data: []*protocol.ClusterServerInfo{
			{
				Type:        protocol.ClusterServerType_LOGIN,
				Ip:          "188.66.66.33",
				Port:        8888,
				CurrentLoad: 5,
			},
		},
	}
	data, err := proto.Marshal(lres)
	t.Log(strconv.Itoa(len(data)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		rdata := &protocol.ClusterResponse{}
		err = proto.Unmarshal(buf[:n], rdata)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("接收数据")
		t.Log(rdata.Status)
		t.Log("测试成功")
		break
	}
}

func Test_RequestRegistDbServer(t *testing.T) {
	service := "localhost:9200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("连接服务器")

	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_REG_SERVER,
		Data: []*protocol.ClusterServerInfo{
			{
				Type: protocol.ClusterServerType_DB,
				Ip:   "188.66.66.133",
				Port: 8888,
			},
		},
	}
	data, err := proto.Marshal(lres)
	t.Log(strconv.Itoa(len(data)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		rdata := &protocol.ClusterResponse{}
		err = proto.Unmarshal(buf[:n], rdata)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("接收数据")
		t.Log(rdata.Status)
		t.Log("测试成功")
		break
	}
}

func Test_RequestGetLoginServers(t *testing.T) {
	service := "localhost:9200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("连接服务器")

	lres := &protocol.ClusterRequest{
		Act: protocol.ClusterActionType_GET_SERVERS,
		ServerType: []protocol.ClusterServerType{
			protocol.ClusterServerType_LOGIN,
			protocol.ClusterServerType_DB,
		},
	}
	data, err := proto.Marshal(lres)
	t.Log(strconv.Itoa(len(data)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		t.Fatal(err)
	}
	// for {
	// 	buf := make([]byte, 1024)
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	t.Log("接收数据")
	// 	t.Log(string(buf[:n]))
	// 	t.Log("测试成功")
	// 	break
	// }
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		rdata := &protocol.ClusterResponse{}
		err = proto.Unmarshal(buf[:n], rdata)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("接收数据")
		t.Log(rdata.Status)
		t.Log(len(rdata.GetData()))
		for _, d := range rdata.GetData() {
			t.Log(d.Ip)
			t.Log(d.LastHeartbeatTime)
		}
		t.Log("测试成功")
		break
	}
}
