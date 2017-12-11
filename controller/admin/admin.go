package admin

import "github.com/kataras/iris"

func Authentication(ctx iris.Context) {
	if true {
		ctx.Next()
	}
}