package main

import (
	"common"
	"common/clog"
	"fmt"
)

const PNAME = "cdrgen"

var log = clog.GetLogger()

func main() {
	fmt.Println("START CDR GENERATOR...")
	log.Info("START CDR GENERATOR...")

	initialize()

	log.Debugf("debug %s", "DEBUG MESSAGE")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}

func initialize() {
	common.ReadConfigFile(PNAME)
	conf := common.GetConfig()
	clog.InitWith(PNAME, conf.Logname, conf.Logdir, conf.Loglevel)

}

// func makeRandomUdr {

// }
