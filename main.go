package main

import (
	"github.com/goes/config"
	"github.com/goes/logger"
	"github.com/goes/model"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"time"
	"github.com/goes/route"
	"strconv"
)

func init() {
	logger.Log(config.DBConfig.Charset)
	logger.Log(config.ServerConfig)
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		logger.Error("open db connect error.")
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	// global
	model.DB = db
}

func main() {
	// 初始化
	app := iris.New(iris.Configuration{
		Gzip: true,
		Charset: "UTF-8",
	})

	// 测试模式
	if config.ServerConfig.Env == model.DevelopmentMode {
		app.Adapt(iris.DevLogger())
	}

	app.Adapt(sessions.New(sessions.Config{
		Cookie:config.ServerConfig.SessionID,
		Expires: time.Minute * 20,
	}))

	app.Adapt(httprouter.New())

	route.Route(app)

	app.OnError(iris.StatusNotFound, func(context *iris.Context) {
		context.JSON(iris.StatusOK, iris.Map{
			"errCode": model.NotFound,
			"message": "Not Found",
			"data": iris.Map{},
		})
	})

	app.OnError(500, func(context *iris.Context) {
		context.JSON(iris.StatusInternalServerError, iris.Map{
			"errCode": model.ERROR,
			"message": "error",
			"data": iris.Map{},
		})
	})

	app.Listen(":" + strconv.Itoa(config.ServerConfig.Port))
}