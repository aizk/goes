package model

import "time"

type Reply struct {
	ID     uint `gorm:"AUTO_INCREMENT;primary_key"`
	UserId uint
	// CommentId uint
	Discuss   Discuss `gorm:"ForeignKey:ProfileID"`
	Reply     string
	Display   bool `gorm:"not null"`
	Describe  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
