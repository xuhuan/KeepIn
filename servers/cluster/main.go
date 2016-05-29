package main

import (
	"github.com/codegangsta/cli"
	// "github.com/golang/protobuf/proto"
	// "github.com/xuhuan/keepin/protocol"
	"github.com/xuhuan/keepin/servers/cluster/cmd"
	"github.com/xuhuan/keepin/utils"
	// "log"
	// "net"
	"os"
	"runtime"
	// "time"
)

const APP_VER = "0.0.1"

var L = utils.L

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "keepin"
	app.Usage = "im server cluster"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdServe,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
	// cmd.Run()
}
