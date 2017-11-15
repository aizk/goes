package main

import (
	"discuss/config"
	"discuss/logger"
	"discuss/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func Log(v ...interface{}) {
	fmt.Println("log:", v)
}

func createDBTable() {

}

func main() {
	Log(config.DBConfig.Charset)
	Log(config.ServerConfig)
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		logger.Error("open db connect error.")
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	model.DB = db
}
