package utils

import (
	"github.com/astaxie/beego/logs"
)

var L = logs.NewLogger(1e3)

func init() {
	L.SetLogger("console", `{"level":7}`)
	L.EnableFuncCallDepth(true)
	L.Async()
}
