package category

import (
	"github.com/kataras/iris"
	"github.com/goes/config"
	"github.com/goes/model"
	"github.com/goes/controller/common"
	"strings"
	"strconv"
	"unicode/utf8"
)

func Save(ctx iris.Context, edit bool) {

	minOrder := config.ServerConfig.MinOrder
	maxOrder := config.ServerConfig.MaxOrder

	var category model.Category

	// 解析参数
	if err := ctx.ReadJSON(&category); err != nil {
		common.SendErrorJSON("参数无效", ctx)
		return
	}

	category.Name = strings.TrimSpace(category.Name)
	if category.Name == "" {
		common.SendErrorJSON("分类名称不能为空", ctx)
		return
	}

	if category.Sequence < minOrder || category.Sequence > maxOrder {
		msg := "分类的排序要在" + strconv.Itoa(minOrder) + "到" + strconv.Itoa(maxOrder) + "之间"
		common.SendErrorJSON(msg, ctx)
		return
	}

	if utf8.RuneCountInString(category.Name) > config.ServerConfig.MaxNameLength {
		msg := "分类名称不能超过" + strconv.Itoa(config.ServerConfig.MaxNameLength) + "个字符"
		common.SendErrorJSON(msg, ctx)
		return
	}

	if category.Status != model.CategoryStatusOpen && category.Status != model.CategoryStatusClose {
		common.SendErrorJSON("无效的 status，必须为1 或者 2", ctx)
		return
	}

	if category.ParentID != 0 {
		var parent model.Category
		if err := model.DB.First(&parent, category.ParentID).Error; err != nil {
			common.SendErrorJSON("无效的父分类", ctx)
			return
		}
	}

	var update model.Category
	if edit {
		// 更新分类
		if err := model.DB.First(&update, category.ID).Error; err != nil {
			common.SendErrorJSON("无效的分类ID", ctx)
			return
		} else {
			// 更新分类
			update.Name = category.Name
			update.Sequence = category.Sequence
			update.ParentID = category.ParentID
			update.Status = category.Status
			if err := model.DB.Save(&update).Error; err != nil {
				common.SendErrorJSON("更新分类失败", ctx)
				return
			}
		}
	} else {
		// 新建分类
		if err := model.DB.Create(&category).Error; err != nil {
			common.SendErrorJSON("创建分类失败", ctx)
			return
		}
	}

	var categoryJSON model.Category

	if edit {
		// 更新分类
		categoryJSON = update
	} else {
		// 新建分类
		categoryJSON = category
	}

	ctx.JSON(iris.Map{
		"errCode": model.SUCCESS,
		"message": "success",
		"data": iris.Map{
			"category": categoryJSON,
		},
	})
	return
}

func Create(ctx *iris.Context) {
	Save(ctx, false)
}

func Update(ctx *iris.Context) {
	Save(ctx, true)
}

// 获取分类信息
func Info(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		common.SendErrorJSON("错误的 ID 类型", ctx)
		return
	}

	var category model.Category
	if err := model.DB.First(&category, id).Error; err != nil {
	    common.SendErrorJSON("查询ID失败", ctx)
	    return
	}

	ctx.JSON(iris.Map{
		"errCode": model.SUCCESS,
		"message": "success",
		"data": iris.Map{
			"category": category,
		},
	})
}

// 更新分类状态 id status

// 获取 status = 1 的可用分类列表

// 获取所有分类
func FetchAllCategory(ctx iris.Context) {
	var categories []model.Category
	pageNum, err := strconv.Atoi(ctx.FormValue("pageNum"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	// 默认为按照创建时间降序排列
	orderString := "created_at"
	if ctx.FormValue("asc") == "1" {
		orderString += " asc"
	} else {
		orderString += " desc"
	}

	offset := (pageNum - 1) * config.ServerConfig.PageSize
	if err := model.DB.Offset(offset).Limit(config.ServerConfig.PageSize).Order(orderString).Find(&categories).Error; err != nil {
		common.SendErrorJSON("查询失败", ctx)
		return
	}

	ctx.JSON(iris.Map{
		"errCode" : model.SUCCESS,
		"message"   : "success",
		"data"  : iris.Map{
			"categories": categories,
		},
	})
}