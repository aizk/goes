package category

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/goes/config"
	"github.com/goes/model"
	"github.com/goes/controller/common"
	"strings"
	"strconv"
)

func Save(ctx *iris.Context, edit bool) {

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

	ctx.JSON(iris.StatusOK, iris.Map{
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

