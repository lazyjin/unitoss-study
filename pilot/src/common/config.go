package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Logdir   string
	Loglevel string
	Logname  string
}

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

func CheckErrPanic(e error) {
	if e != nil {
		fmt.Printf("%v\n", e)
		panic(e)
	}
}
