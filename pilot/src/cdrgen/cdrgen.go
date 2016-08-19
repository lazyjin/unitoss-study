package main

import (
	"common"
	"common/clog"
	"fmt"
	"math/rand"
	"time"
	"udr"
)

const PNAME = "cdrgen"

var (
	log       = clog.GetLogger()
	rabbitMgr = common.NewRabbitManager()
)

func main() {
	log.Info("START CDR GENERATOR...")

	initialize()

	randUdr := makeRandomUdr()

	fmt.Printf("random UDR %v\n", randUdr)

	jsonUdr, err := randUdr.ConvToJsonStr()
	if err != nil {
		log.Errorf("Udr to Json failed: UDR: [%v] JSON: [%v]", randUdr, err)
	}

	fmt.Printf("random JSON UDR %v\n", jsonUdr)

	err = rabbitMgr.PublishToQueue(jsonUdr)
	if err != nil {
		log.Errorf("UDR message is not send: %v", err)
	}

	// log.Debugf("debug %s", "DEBUG MESSAGE")
	// log.Info("info")
	// log.Notice("notice")
	// log.Warning("warning")
	// log.Error("err")
	// log.Critical("crit")
}

func initialize() {
	common.ReadConfigFile(PNAME)
	conf := common.GetConfig()
	log.Debugf("Config: %v", conf)
	// init log
	clog.InitWith(PNAME, conf.Logname, conf.Logdir, conf.Loglevel)

	// rabbit publisher connect
	rabbitMgr.ConnectRabbit(
		conf.Rabbithost,
		conf.Rabbitport,
		conf.Rabbituser,
		conf.Rabbitpw)
	rabbitMgr.TaskQueueDeclare("task_queue")
}

func makeRandomUdr() udr.UdrRaw {
	tmpUdr := udr.GetEmptyUdrRaw()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make random EUI && byte count
	randEui := r.Uint32()%udr.EUI_BASE + udr.EUI_BASE
	randByte := r.Uint32() % 10 * 100

	// make time fields
	now := time.Now()
	start := fmt.Sprintf(
		common.TIME_FMT,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond()/1000000000)

	d, err := time.ParseDuration("5s")
	if err != nil {
		log.Panic(err)
	}

	then := now.Add(d)
	end := fmt.Sprintf(
		common.TIME_FMT,
		then.Year(),
		then.Month(),
		then.Day(),
		then.Hour(),
		then.Minute(),
		then.Second(),
		then.Nanosecond()/1000000000)

	tmpUdr.SetUdrRaw(randEui, start, end, randByte, "")

	return tmpUdr
}
