package main

import (
	"log"
	"os"
)

func main() {
	INFO.Println("hahaha")
}

var (
	INFO *log.Logger
)

func init() {
	filename := "info.log"
	logFile, err := os.Create(filename)
	if err != nil {
		log.Fatalln("create log file err", err)
	}
	INFO = log.New(logFile, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
}
