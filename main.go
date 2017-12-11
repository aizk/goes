package main

import (
	"github.com/goes/config"
	. "github.com/goes/logger"
	"github.com/goes/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"os"
	//"github.com/kataras/iris/sessions"
	"github.com/goes/route"
	"github.com/kataras/iris/middleware/logger"
	"strconv"
)

func init() {
	Log(config.DBConfig.Charset)
	Log(config.ServerConfig)
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		Error("open db connect error.")
		os.Exit(-1)
	}
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	// global
	model.DB = db
}

func main() {

	app := iris.New()

	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	app.Use(logger.New())

	route.Route(app)

	// 测试模式
	//if config.ServerConfig.Env == model.DevelopmentMode {
	//	app.Adapt(iris.DevLogger())
	//}

	//app.Adapt(sessions.New(sessions.Config{
	//	Cookie:config.ServerConfig.SessionID,
	//	Expires: time.Minute * 20,
	//}))

	//app.Adapt(httprouter.New())

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"err":  model.NotFound,
			"msg":  "Not Found",
			"data": iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"err":     model.ERROR,
			"message": "error",
			"data":    iris.Map{},
		})
	})

	address := iris.Addr(":" + strconv.Itoa(config.ServerConfig.Port))

	if config.ServerConfig.Env == model.DevelopmentMode {
		app.Run(address)
	} else {
		app.Run(address, iris.WithoutVersionChecker)
	}
}
