package model

import (
	"time"
)

type Discuss struct {
	ID        uint `gorm:"not null;AUTO_INCREMENT"`
	Fk        uint
	FkSub     uint
	UserID    uint `gorm:"not null"`
	Comment   string `gorm:"not null"`
	Type      uint `gorm:"not null"`
	TypeSub   uint
	Display   bool
	Describe  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

const (
	// 题库评论
	SubjectComment = 1
	// 题库笔记
	SubjectNote = 2
	// 题库纠错
	SubjectError = 3
	// 题目视频
	SubjectVideo = 4
	// 网校视频
	OnlineSchoolVideo = 5
	// 建议
	Suggest = 6
)
