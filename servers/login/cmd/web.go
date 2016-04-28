package cmd

import (
	"github.com/astaxie/beego"
	"github.com/codegangsta/cli"
	_ "github.com/xuhuan/keepin/servers/login/docs"
	_ "github.com/xuhuan/keepin/servers/login/routers"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "start login server",
	Description: `login server for client`,
	Action:      runWeb,
	Flags: []cli.Flag{
		stringFlag("config, c", "conf/app.conf", "configuration file path"),
	},
}

func runWeb(ctx *cli.Context) {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func Run() {
	beego.Run()
}
