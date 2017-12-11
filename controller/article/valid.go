package article

import (
	"github.com/goes/config"
	"github.com/goes/controller/common"
	"github.com/goes/model"
	"github.com/kataras/iris"
	"strconv"
	"strings"
	"unicode/utf8"
)

func articleValid(article *model.Article, ctx iris.Context) {
	if article.Name == "" {
		common.SendErrorJSON("文章名字不能为空", ctx)
		return
	}
	if utf8.RuneCountInString(article.Name) > config.ServerConfig.MaxNameLength {
		common.SendErrorJSON("文章名称长度超过最大长度 " + strconv.Itoa(config.ServerConfig.MaxNameLength), ctx)
		return
	}
	// 文章分类
	if article.Categories == nil || len(article.Categories) <= 0 {
		common.SendErrorJSON("请选择文章分类", ctx)
		return
	}

	if len(article.Categories) > config.ServerConfig.MaxArticleCateCount {
		msg := "文章最多属于" + strconv.Itoa(config.ServerConfig.MaxArticleCateCount) + "个版块"
		common.SendErrorJSON(msg, ctx)
		return
	}

	// 验证分类是否有效
	for i := 0; i < len(article.Categories); i++ {
		var category model.Category
		if err := model.DB.First(&category, article.Categories[i].ID).Error; err != nil {
			common.SendErrorJSON("无效的分类id", ctx)
			return
		}
		article.Categories[i] = category
	}

	article.Name = strings.TrimSpace(article.Name)
	article.Content = strings.TrimSpace(article.Content)
}
