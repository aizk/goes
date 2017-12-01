package article

import (
	"github.com/goes/model"
	"strconv"
	"github.com/kataras/iris"
	"github.com/goes/config"
	"time"
	"github.com/goes/controller/common"
)

func ArticleList(ctx iris.Context)  {
	var articles []model.Article
	var pageNum int
	var err error

	if pageNum, err = strconv.Atoi(ctx.FormValue("pageNum")); err != nil {
		// 读取失败，默认为第一页
		pageNum = 1
		err = nil
	}

	if pageNum < 1 {
		pageNum = 1
	}

	pageSize := config.ServerConfig.PageSize

	offset := (pageNum - 1) * pageSize

	var startTime string
	var endTime string

	// 时间查询
	if startAt, err := strconv.Atoi(ctx.FormValue("startAt")); err != nil {
		// 1970 年的时间
		startTime = time.Unix(0, 0).Format("2006-01-02 15:04:05")
	} else {
		startTime = time.Unix(int64(startAt/1000), 0).Format("2006-01-02 15:04:05")
	}

	if endAt, err := strconv.Atoi(ctx.FormValue("endAt")); err != nil {
		// 当前时间
		endTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		endTime = time.Unix(int64(endAt / 1000), 0).Format("2006-01-02 15:04:05")
	}

	// 默认按创建时间 降序排列
	order := "created_at"
	var orderASC string
	if ctx.FormValue("asc") == "1" {
		orderASC = "ASC"
	} else {
		orderASC = "DESC"
	}

	cateID := ctx.FormValue("cateId")
	var categoryID int
	// Atoi 转换空字符串为 0
	if categoryID, err = strconv.Atoi(cateID); err != nil {
		common.SendErrorJSON("错误的分类 ID", ctx)
		return
	}


}