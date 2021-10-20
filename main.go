package main

import (
	glua "golua/conf"
	logger "golua/logger"
)

const VERSION = "v1.0.0"

func main() {
	// logo()
	println(glua.GluaConf.Http_server.Port)
}

func logo() {
	logger.INFO.Println("=================================")
	logger.INFO.Printf("========= golua %s ==========\n", VERSION)
	logger.INFO.Println("=================================")

}
