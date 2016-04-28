package client

import (
	"github.com/astaxie/beego/config"
	// "strconv"
)

func InitDb() {

	serverConf, err := config.NewConfig("json", "conf/server.json")
	if err != nil {
		L.Critical("配置文件不存在")
	}
	L.Debug("连接数据服务器")
	L.Debug("登录连接超时时间：%s 毫秒", serverConf.String("LoginResponseTimeout"))
	ss := serverConf.Strings("DbServers")
	if len(ss) > 0 {
		L.Debug("解析成功")
	}
	// L.Debug(strconv.FormatInt(len(ss)))
	// if err != nil {
	// 	L.Debug("解析成功")
	// } else {
	// 	L.Debug("解析错误")
	// }
	// L.Debug(string(servers["0"]["ip"]))

}
