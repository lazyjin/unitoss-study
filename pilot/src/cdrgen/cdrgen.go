package main

import (
	"common"
	"common/clog"
	"common/rabbit"
	"common/redismgr"
	"common/udr"
	"flag"
	"reflect"
	"time"
)

const PNAME = "cdrgen"

var (
	log       = clog.GetLogger()
	rabbitMgr = rabbit.NewRabbitManager()
	redis     = redismgr.GetRedisCluster()
)

func main() {
	isDaemon := flag.Bool("d", false, "Daemon mode. subscribe request from web")
	flag.Parse()

	log.Info("START CDR GENERATOR...")

	defer rabbitMgr.CloseChanRabbit()
	defer rabbitMgr.CloseConnRabbit()

	if *isDaemon {
		log.Info("Running Daemon mode...")
		setRecvQueue()
		runDaemon()

	} else {
		procCh := make(chan bool)
		processUdrGen(1, udr.MakeRandomUdr, procCh)
		<-procCh
	}
}

func init() {
	common.ReadConfigFile(PNAME)
	conf := common.GetConfig()
	// init log
	clog.InitWith(PNAME, conf.Logname, conf.Logdir, conf.Loglevel)
	log.Info("CDRGEN Initializing...")

	// rabbit publisher connect
	rabbitMgr.ConnectRabbit(
		conf.Rabbithost,
		conf.Rabbitport,
		conf.Rabbituser,
		conf.Rabbitpw)
	rabbitMgr.UdrSendQueueDeclare(conf.Udrqueue)

	// redis cluster connect
	redismgr.ConnectRedisCluster(conf.Redisclusters)
}

func setRecvQueue() {
	conf := common.GetConfig()
	rabbitMgr.ReqRecvQueueDeclare(conf.Reqreciever)
}

func runDaemon() {
	msgs, err := rabbitMgr.ConsumeQueue()
	if err != nil {
		log.Panicf("Fail to Consume Queue...[%v]", err)
	}

	// make blocking for get rabbitmq msessages
	forever := make(chan bool)

	// messgae processing main loop
	go processUdrReq(msgs)

	log.Info("Waiting for UDR request messages...")
	<-forever
}

func processUdrReq(msgs rabbit.QueMsg) {
	procCh := make(chan bool)

	for d := range msgs {
		log.Infof("Received a message: %s %v", d.Body, reflect.TypeOf(d.Body))

		reqMsg, err := common.UdrReqMsgParse(d.Body)
		if err != nil {
			log.Errorf("Fail to parse UDR request message => %v", err)
			d.Reject(false)
			continue
		}

		var udrFunc func() (udr.UdrRaw, error)

		switch reqMsg.ErrorType {
		case common.NORMAL:
			log.Info("Generate Normal UDR...")
			udrFunc = udr.MakeRandomUdr
		case common.TIME_ERR:
			log.Info("Generate Time Error UDR...")
			udrFunc = udr.MakeTimeErrUdr
		case common.EUI_ERR:
			log.Info("Generate EUI Error UDR...")
			udrFunc = udr.MakeEuiErrUdr
		case common.FMT_ERR:
			log.Info("Generate Format Error UDR...")
			udrFunc = udr.MakeFmtErrUdr
		}

		go processUdrGen(reqMsg.Count, udrFunc, procCh)
		go rabbit.ResponseAck(d, procCh)
	}
}

func processUdrGen(udrCnt int, procFunc func() (udr.UdrRaw, error), procCh chan bool) {
	start := time.Now()

	for i := 0; i < udrCnt; i++ {
		randUdr, err := procFunc()
		if err != nil {
			log.Errorf("Fail to make random udr: [%v]", err)
			procCh <- false
			break
		}

		jsonUdr, err := randUdr.ConvToJsonStr()
		if err != nil {
			log.Errorf("Udr to Json failed: UDR: [%v] JSON: [%v]", randUdr, err)
			procCh <- false
			break
		}

		log.Debugf("random JSON UDR %v\n", jsonUdr)

		err = rabbitMgr.PublishToQueue(jsonUdr)
		if err != nil {
			log.Errorf("UDR message is not send: %v", err)
			procCh <- false
			break
		}
	}

	procCh <- true

	elapsed := time.Since(start)
	log.Infof("Generate %d UDR takes %s...", udrCnt, elapsed)
}
