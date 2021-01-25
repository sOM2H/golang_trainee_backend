package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/model"
)

type userUpdateRequest struct {
	User struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func newUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

func (r *userUpdateRequest) populate(u *model.User) {
	r.User.Email = u.Email
	r.User.Password = u.Password
}

func (r *userUpdateRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.User.Email
	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	return nil
}

type userRegisterRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type postCreateRequest struct {
	Post struct {
		Title string `json:"title" validate:"required"`
		Body  string `json:"body" validate:"required"`
	} `json:"post"`
}

func (r *postCreateRequest) bind(c echo.Context, a *model.Post) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Post.Title
	a.Body = r.Post.Body
	return nil
}

type postUpdateRequest struct {
	Post struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"post"`
}

func (r *postUpdateRequest) populate(a *model.Post) {
	r.Post.Title = a.Title
	r.Post.Body = a.Body
}

func (r *postUpdateRequest) bind(c echo.Context, a *model.Post) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Post.Title
	a.Body = r.Post.Body
	return nil
}

type createCommentRequest struct {
	Comment struct {
		Body   string `json:"body" validate:"required"`
		PostID uint   `json:"post_id"`
	} `json:"comment"`
}

func (r *createCommentRequest) bind(c echo.Context, cm *model.Comment) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	postID, _ := strconv.Atoi(c.FormValue("post_id"))

	cm.Body = r.Comment.Body
	cm.PostID = uint(postID)
	cm.AuthorID = userIDFromToken(c)
	return nil
}

type commentUpdateRequest struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"post"`
}

func (r *commentUpdateRequest) populate(a *model.Comment) {
	r.Comment.Body = a.Body
}

func (r *commentUpdateRequest) bind(c echo.Context, a *model.Comment) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Body = r.Comment.Body
	return nil
}
