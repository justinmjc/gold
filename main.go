package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "gold/routers"
)

func main() {
	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m",
		WhenFormat: "2006-01-02",
	}
	logs.RegisterFormatter("pattern", f)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)

	_ = logs.SetGlobalFormatter("pattern")


	logs.Info("hello, world")
	beego.Run()
}

