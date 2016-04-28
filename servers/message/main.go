package main

import (
	_ "github.com/xuhuan/keepin/servers/message/docs"
	_ "github.com/xuhuan/keepin/servers/message/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
