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

var dbPoolR = make(map[string]db_client)

var dbPoolW = make(map[string]db_client)

/**
read client has no transaction
*/
func GetReadClient(dbName string) *db_client {
	if len(dbName) == 0 {
		logger.ERR.Fatalln("dbName is empty")
	}
	dbClient, ok := dbPool[dbName]
	if ok {
		return &dbClient
	}
	newDbClient := createClient(dbName)
	dbPool[dbName] = newDbClient
	return newDbClient
}

func createClient(dbName string, isTransaction bool) *db_client {
	switch dbName {
	case "mysql":
		sql.Open("mysql", dbConf.Mysql.Dsn)
	case "oracle":
		// todo
	default:
		logger.ERR.Fatalf("db not support, dbName = %s\n", dbName)
	}
	return nil
}

func init() {
	decoder := util.ReadJson(db_conf_name)
	err := decoder.Decode(&dbConf)
	if err != nil {
		logger.ERR.Fatalf("decode %s error, err = %s\n", db_conf_name, err.Error())
	}
}
