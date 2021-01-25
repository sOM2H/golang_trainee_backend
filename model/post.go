package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Body     string
	Author   User
	AuthorID uint
	Comments []Comment
}
