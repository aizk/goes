package route

import (
	"gopkg.in/kataras/iris.v6"
)

const Prefix  = "goes"

// Route 路由
func Route(app *iris.Framework) {
	apiPrefix := Prefix

	router := app.Party(apiPrefix)

	router.Get("/categories", )
}