package post

import (
	"github.com/sOM2H/golang_trainee_backend/model"
)

type Store interface {
	GetPostByID(uint) (*model.Post, error)
	CreatePost(*model.Post) error
	UpdatePost(*model.Post) error
	DeletePost(*model.Post) error
	ListPost(offset, limit int) ([]model.Post, int, error)
}
