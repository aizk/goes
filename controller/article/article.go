package article

import (
	"github.com/goes/config"
	"github.com/goes/controller/common"
	"github.com/goes/model"
	"github.com/kataras/iris"
	"strconv"
	"time"
)

func articleList(getAll bool, ctx iris.Context) {
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

	//var startTime string
	//var endTime string
	//
	//// 时间查询
	//if startAt, err := strconv.Atoi(ctx.FormValue("startAt")); err != nil {
	//	// 1970 年的时间
	//	startTime = time.Unix(0, 0).Format("2006-01-02 15:04:05")
	//} else {
	//	startTime = time.Unix(int64(startAt/1000), 0).Format("2006-01-02 15:04:05")
	//}
	//
	//if endAt, err := strconv.Atoi(ctx.FormValue("endAt")); err != nil {
	//	// 当前时间
	//	endTime = time.Now().Format("2006-01-02 15:04:05")
	//} else {
	//	endTime = time.Unix(int64(endAt/1000), 0).Format("2006-01-02 15:04:05")
	//}

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
		common.SendErrorJSON("错误的分类 ID 类型", ctx)
		return
	}

	if categoryID > 0 {
		var category model.Category
		if model.DB.First(&category, categoryID).Error != nil {
			common.SendErrorJSON("查询失败", ctx)
			return
		}
		var sql = `SELECT distinct(article.id), article.name, article.browse_count, article.status, article.created_at, article.update_at
				FROM atricle as a, article_category as ac
				WHERE article.id = article_category.article_id AND article_category.category_id = ?
				ORDER BY ? ?
				LIMIT(?, ?)
				`
		err = model.DB.Exec(sql, categoryID, order, orderASC, offset, pageSize).Scan(&articles).Error
		if err != nil {
			common.SendErrorJSON("Exec 执行 SQL 出错", ctx)
			return
		}
		for i := 0; i < len(articles); i++ {
			articles[i].Categories = []model.Category{category} // List
		}
	} else {
		// 没有 categoryId 则获取所有数据
		orderString := order + " " + orderASC
		// 后台获取所有数据
		if getAll {
			err = model.DB.Offset(offset).Limit(pageSize).Order(orderString).Find(&articles).Error
		} else {
			err = model.DB.Where("status = 1 OR status = 2").Offset(offset).Limit(pageSize).Order(orderString).Find(&articles).Error
		}
		if err != nil {
			common.SendErrorJSON("查询所有文章失败", ctx)
			return
		}
		for i := 0; i < len(articles); i++ {
			if err = model.DB.Model(&articles[i]).Related(&articles[i].Categories, "categories").Error; err != nil {
				common.SendErrorJSON("Related Error", ctx)
				return
			}
		}
	}
	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": iris.Map{
			"articles": articles,
			"pageNum":  pageNum,
			"pageSize": pageSize,
		},
	})
}

func List(ctx iris.Context) {
	articleList(false, ctx)
}

// 后台获取所有文章接口
func AllList(ctx iris.Context) {
	articleList(true, ctx)
}

func Create(ctx iris.Context) {
	var article model.Article
	if err := ctx.ReadJSON(&article); err != nil {
		common.SendErrorJSON("无效参数", ctx)
		return
	}
	articleValid(&article, ctx)
	article.BrowseCount = 0
	// 审核中
	article.Status = model.ArticleVerifying
	if err := model.DB.Create(&article).Error; err != nil {
		common.SendErrorJSON("创建文章失败", ctx)
		return
	}
	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": article,
	})
}

func Update(ctx iris.Context) {
	var article model.Article
	if err := ctx.ReadJSON(&article); err != nil {
		common.SendErrorJSON("无效参数", ctx)
		return
	}
	articleValid(&article, ctx)

	var databaseArticle model.Article
	if err := model.DB.First(&databaseArticle, article.ID).Error; err != nil {
	    common.SendErrorJSON("无效的文章 ID", ctx)
	    return
	}
	article.BrowseCount = databaseArticle.BrowseCount
	article.CreatedAt = databaseArticle.CreatedAt
	article.Status = databaseArticle.Status
	article.UpdatedAt = time.Now()
	if err := model.DB.Save(&article).Error; err != nil {
	    common.SendErrorJSON("更新文章失败", ctx)
	    return
	}
	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": article,
	})
}

func Read(ctx iris.Context) {
	var id int
	id, err := ctx.URLParamInt("id")

	if err != nil {
		common.SendErrorJSON("错误的文章id", ctx)
		return
	}
	var article model.Article
	if err := model.DB.First(&article, id).Error; err != nil {
		common.SendErrorJSON("获取文章失败", ctx)
		return
	}
	// 填充分类信息
	if err := model.DB.Model(&article).Related(&article.Categories, "categories").Error; err != nil {
	    common.SendErrorJSON("填充分类信息失败", ctx)
	    return
	}
	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": article,
	})
}

// 更新文章状态
func UpdateStatus(ctx iris.Context) {
	return
}