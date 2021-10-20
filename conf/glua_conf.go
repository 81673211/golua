package conf

import (
	"encoding/json"
	"golua/logger"
	"os"
)

const glua_conf_name = "conf/glua.json"

type http_server struct {
	Port    int
	Enable  bool
	Threads int
}

type glua_conf struct {
	Http_server http_server
}

var GluaConf = glua_conf{}

func init() {
	file, err := os.Open(glua_conf_name)
	if err != nil {
		logger.ERR.Fatalln("open glua.json error", err)
	}
	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&GluaConf)
	if decodeErr != nil {
		logger.ERR.Fatalln("decode glua.json error", decodeErr)
	}
}
