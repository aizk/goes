package model

import "time"

// 文章分类
type Category struct {
	ID        uint       `gorm:"primary_key`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Name      string
	Sequence  int
	ParentID  int
	Status    int
}

// 开启
const CategoryStatusOpen = 1

// 关闭
const CategoryStatusClose = 2
