package config

import (
	"github.com/goes/logger"
	"github.com/goes/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		logger.Error("ReadFile: ", err.Error())
		os.Exit(-1)
	}
	// 去除注释
	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		logger.Log("json parse fail", err.Error())
		os.Exit(-1)
	}
}

type dbConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Charset      string
	Host         string
	Port         int
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

var DBConfig dbConfig

func initDB() {
	utils.SetObjectByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User,
		DBConfig.Password,
		DBConfig.Host,
		DBConfig.Port,
		DBConfig.Database,
		DBConfig.Charset)
	DBConfig.URL = url
}

type serverConfig struct {
	Env string
	SessionID string
	Port int
	PageSize int
	MaxPageSize int
	MinPageSize int
	MinOrder int
	MaxOrder int
	MaxNameLength int
	MaxContentLength int
}

var ServerConfig serverConfig

func initServer() {
	utils.SetObjectByJSON(&ServerConfig, jsonData["server"].(map[string]interface{}))
}

func init() {
	initJSON()
	initDB()
	initServer()
}