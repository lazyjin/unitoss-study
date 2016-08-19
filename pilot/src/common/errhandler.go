package common

import (
	"common/clog"
)

var log = clog.GetLogger()

func CheckErrPanic(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func CheckErr(e error) {
	if e != nil {
		log.Error(e)
	}
}
