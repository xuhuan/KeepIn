package main

import (
	// "github.com/codegangsta/cli"
	"github.com/xuhuan/KeepIn_Server/servers/login/client"
	"github.com/xuhuan/KeepIn_Server/servers/login/cmd"
	"github.com/xuhuan/KeepIn_Server/servers/login/utils"
	// "os"
	"github.com/googollee/go-socket.io"
	"runtime"
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
	client.Init()
	// cmd.Run()

}
