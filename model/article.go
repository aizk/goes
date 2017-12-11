package model

import "time"

// Article 与 Category 属于多对多关系，指定使用哪张表连接
// 查询文章的所有分类：                         Categories 是关系中源 Article 内的字段名
// db.Model(&article).Related(&categories, "Categories")
// 指定外键 ForeignKey:Id
// 指定关联外键 AssociationForeignKey:Id
type Article struct {
	ID          uint `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAt"`
	Name        string `json:"name"`
	BrowseCount int `json:"browseCount"`
	Status      int `json:"status"`
	Content     string `json:"content"`
	Categories  []Category `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:Id" json:"categories"`
	Comments []Comment `json:"comments"`
}

const (
	// ArticleVerifying 审核中
	ArticleVerifying = 1

	// ArticleVerifySuccess 审核通过
	ArticleVerifySuccess = 2

	// ArticleVerifyFail 审核未通过
	ArticleVerifyFail = 3
)
