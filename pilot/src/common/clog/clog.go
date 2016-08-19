package clog

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)

// Log levels.
const (
	CRITICAL logging.Level = logging.CRITICAL
	ERROR                  = logging.ERROR
	WARNING                = logging.WARNING
	NOTICE                 = logging.NOTICE
	INFO                   = logging.INFO
	DEBUG                  = logging.DEBUG
)

var (
	log logging.Logger

	logFilename string
	logPath     string
)

var format = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05.999} %{shortfunc} %{shortfile} [%{level:.4s}] %{message}`,
)

func Init() {
	logging.MustGetLogger("default")
	defaultBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(defaultBackend, format)

	logging.SetBackend(backendFormatter)
}

func InitWith(logname string, filename string, path string, level string) {
	logging.MustGetLogger(logname)
	logFilename = filename
	logPath = path

	logFullPath := path + "/" + filename

	f, err := os.OpenFile(logFullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Fail to open file %v\n", logFullPath)
		panic("Fail to open file")
	}

	backend := logging.NewLogBackend(f, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	lev, err := logging.LogLevel(level)
	if err != nil {
		panic(err)
	}

	backendLeveled.SetLevel(lev, logname)

	// Set the backends to be used.
	logging.SetBackend(backendLeveled)
}

func GetLogger() *logging.Logger {
	return &log
}
