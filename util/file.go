package util

import (
	"encoding/json"
	"golua/logger"
	"os"
)

func ReadJson(filename string) *json.Decoder {
	file, err := os.Open(filename)
	if err != nil {
		logger.ERR.Fatalf("open %s error, err = %s\n", filename, err.Error())
	}
	return json.NewDecoder(file)
}
