package admin

import "gopkg.in/kataras/iris.v6"

func Authentication(ctx *iris.Context) {
	if true {
		ctx.Next()
	}
}