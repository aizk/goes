package common

import (
	"github.com/kataras/iris"
	"github.com/goes/model"
)

func SendErrorJSON(message string, ctx iris.Context)  {
	ctx.JSON(iris.Map{
		"errCode": model.ERROR,
		"message": message,
		"data": iris.Map{},
	})
}