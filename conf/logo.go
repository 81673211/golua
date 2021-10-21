package conf

import "golua/logger"

const VERSION = "v1.0.0"

func init() {
	logger.INFO.Println("=================================")
	logger.INFO.Printf("========= golua %s ==========\n", VERSION)
	logger.INFO.Println("=================================")
}
