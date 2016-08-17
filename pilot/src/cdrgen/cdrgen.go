package main

import (
	"common/clog"
	"fmt"
)

var log = clog.GetLogger()

func main() {
	fmt.Println("START CDR GENERATOR...")

	initialize()

	log.Debugf("debug %s", "DEBUG MESSAGE")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}

func initialize() {
	clog.InitWith("cdrgen", "app_cdrgen.log", ".", clog.DEBUG)
}
