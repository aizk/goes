package route

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/goes/controller/category"
	"github.com/goes/controller/admin"
)

const Prefix  = "goes"

// Route 路由
func Route(app *iris.Framework) {
	apiPrefix := Prefix

	router := app.Party(apiPrefix)
	{
		router.Get("/categories", nil)
	}

	adminRouter := app.Party(apiPrefix + "/admin", admin.Authentication)
	{
		adminRouter.Post("/category/create", category.Create)
		adminRouter.Post("/category/update", category.Update)
	}
}