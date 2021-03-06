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
	Rabbituser string
	Rabbitpw   string

	// queue information
	Udrqueue    string
	Reqreciever string

	// redis config
	Redisclusters []string

	// web template
	Templatedir string
}

var conf Config

func ReadConfigFile(pname string) {
	cfgpath := os.Getenv("CFG_DIR")

	data, err := ioutil.ReadFile(cfgpath + "/" + pname + ".yaml")
	// test code for windows
	// data, err := ioutil.ReadFile(cfgpath + "\\cdrgen.yaml")
	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() *Config {
	return &conf
}
