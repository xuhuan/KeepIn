package main

import (
	_ "github.com/xuhuan/KeepIn_Server/servers/route/docs"
	_ "github.com/xuhuan/KeepIn_Server/servers/route/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
