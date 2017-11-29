package common

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/goes/model"
)

func SendErrorJSON(message string, ctx *iris.Context)  {
	ctx.JSON(iris.StatusOK, iris.Map{
		"errCode": model.ERROR,
		"message": message,
		"data": iris.Map{},
	})
}