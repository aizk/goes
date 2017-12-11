package comment

import (
	"github.com/goes/controller/common"
	"github.com/goes/model"
	"github.com/kataras/iris"
)

func Create(ctx iris.Context) {
	var comment model.Comment
	commentValid(&comment, ctx)
	comment.Status = model.CommentVerifying
	if err := model.DB.Create(&comment).Error; err != nil {
	    common.SendErrorJSON("新建评论失败", ctx)
	    return
	}
	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": comment,
	})
}

func Update(ctx iris.Context) {
	var comment model.Comment
	var updateComment model.Comment
	commentValid(&comment, ctx)

	if err := model.DB.First(&updateComment, comment.Id).Error; err == nil {
	    updateComment.Content = comment.Content
		updateComment.Status = model.CommentVerifying
		if err := model.DB.Save(&updateComment).Error; err != nil {
		    common.SendErrorJSON("修改评论失败", ctx)
		    return
		}
	} else {
		common.SendErrorJSON("无效的评论", ctx)
		return
	}

	ctx.JSON(iris.Map{
		"err": model.SUCCESS,
		"msg": "success",
		"data": iris.Map{
			"comment": updateComment,
		},
	})
}