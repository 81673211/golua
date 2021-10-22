package db

import (
	"database/sql"
	"golua/logger"
	"golua/util"
)

const db_conf_name = "etc/json/db.json"

type db_client interface {
}

type mysql_client struct {
}

type mysql_conf struct {
	Dsn             string
	ConnMaxLifetime int
	MaxOpenConns    int
	MaxIdleConns    int
}

type db_conf struct {
	Mysql mysql_conf
}

var dbConf = db_conf{}

var dbPool = make(map[string]db_client)

/**
read client has no transaction
*/
func GetMysqlReadClient() *db_client {
	sql.Open("mysql", dbConf.Mysql.Dsn)
	return nil
}

func init() {
	decoder := util.ReadJson(db_conf_name)
	err := decoder.Decode(&dbConf)
	if err != nil {
		logger.ERR.Fatalf("decode %s error, err = %s\n", db_conf_name, err.Error())
	}
}
