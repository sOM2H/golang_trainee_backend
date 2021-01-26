package comment

import (
	"github.com/sOM2H/golang_trainee_backend/model"
)

type Store interface {
	GetCommentByID(uint) (*model.Comment, error)
	CreateComment(*model.Comment) error
	UpdateComment(*model.Comment) error
	DeleteComment(*model.Comment) error
	ListCommentByPostID(offset, limit, id int) ([]model.Comment, int, error)
}
