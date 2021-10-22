package conf

import (
	"golua/logger"
	"golua/util"
)

const glua_conf_name = "etc/json/glua.json"

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
	decoder := util.ReadJson(glua_conf_name)
	err := decoder.Decode(&GluaConf)
	if err != nil {
		logger.ERR.Fatalf("decode %s error, err = %s\n", glua_conf_name, err.Error())
	}
}
