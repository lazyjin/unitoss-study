package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	// Log config
	Logdir   string
	Loglevel string
	Logname  string

	// RabbitMQ config
	Rabbithost string
	Rabbitport int
}

const TIME_FMT = "%d%02d%02d%02d%02d%02d%1d"

var conf Config

func ReadConfigFile(pname string) {
	cfgpath := os.Getenv("CFG_DIR")
	// cfgfile := os.Args[0]

	data, err := ioutil.ReadFile(cfgpath + "/" + pname + ".yaml")
	// test code for windows
	// data, err := ioutil.ReadFile(cfgpath + "\\cdrgen.yaml")
	CheckErrPanic(err)

	err = yaml.Unmarshal(data, &conf)
	CheckErrPanic(err)
}

func GetConfig() *Config {
	return &conf
}
