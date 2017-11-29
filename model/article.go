package model

import "time"

type Article struct {
	ID          uint `gorm:"primary_key`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	BrowseCount int
	Status      int
	Content     string
	Categories  []Category
}

const (
	// ArticleVerifying 审核中
	ArticleVerifying      = 1

	// ArticleVerifySuccess 审核通过
	ArticleVerifySuccess  = 2

	// ArticleVerifyFail 审核未通过
	ArticleVerifyFail     = 3
)