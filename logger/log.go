package logger

import (
	"log"
	"os"
)

var (
	DEBUG *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERR   *log.Logger
)

func init() {
	filename := "all.log"

	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("create log file err", err)
	}
	DEBUG = log.New(logFile, "[DEBUG]", log.LstdFlags|log.Llongfile)
	INFO = log.New(logFile, "[INFO]", log.LstdFlags|log.Llongfile)
	WARN = log.New(logFile, "[WARN]", log.LstdFlags|log.Llongfile)
	ERR = log.New(logFile, "[ERR]", log.LstdFlags|log.Llongfile)
}
