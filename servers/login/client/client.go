package client

import (
	"github.com/xuhuan/KeepIn_Server/servers/login/utils"
)

var L = utils.L

func Init() {
	L.Info("初始化连接服务客户端")
	InitDb()
}
