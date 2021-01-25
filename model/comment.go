package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Post     Post
	PostID   uint
	Author   User
	AuthorID uint
	Body     string
}
