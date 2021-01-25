package handler

import (
	"github.com/sOM2H/golang_trainee_backend/controllers/comment"
	"github.com/sOM2H/golang_trainee_backend/controllers/post"
	"github.com/sOM2H/golang_trainee_backend/controllers/user"
)

type Handler struct {
	userStore    user.Store
	postStore    post.Store
	commentStore comment.Store
}

func NewHandler(us user.Store, ps post.Store, cs comment.Store) *Handler {
	return &Handler{
		commentStore: cs,
		userStore:    us,
		postStore:    ps,
	}
}
