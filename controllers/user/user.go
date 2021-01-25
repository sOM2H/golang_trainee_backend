package user

import (
	"github.com/sOM2H/golang_trainee_backend/model"
)

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	Show(uint) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
}
