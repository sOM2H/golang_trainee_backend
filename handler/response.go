package handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/controllers/user"
	"github.com/sOM2H/golang_trainee_backend/model"
	"github.com/sOM2H/golang_trainee_backend/utils"
)

type userResponse struct {
	User struct {
		Email string `json:"email"`
		Token string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type profileResponse struct {
	Profile struct {
	} `json:"profile"`
}

func newProfileResponse(us user.Store, userID uint, u *model.User) *profileResponse {
	r := new(profileResponse)
	return r
}

type postResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AuthorID  uint      `json:"author_id"`
}

type singlePostResponse struct {
	Post *postResponse `json:"post"`
}

type postListResponse struct {
	Posts      []*postResponse `json:"posts"`
	PostsCount int             `json:"postsCount"`
}

func newPostResponse(c echo.Context, a *model.Post) *singlePostResponse {
	ar := new(postResponse)
	ar.AuthorID = a.AuthorID
	ar.ID = a.ID
	ar.Title = a.Title
	ar.Body = a.Body
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt
	return &singlePostResponse{ar}
}

func newPostListResponse(us user.Store, userID uint, posts []model.Post, count int) *postListResponse {
	r := new(postListResponse)
	r.Posts = make([]*postResponse, 0)
	for _, a := range posts {
		ar := new(postResponse)
		ar.ID = a.ID
		ar.Title = a.Title
		ar.Body = a.Body
		ar.AuthorID = a.AuthorID
		ar.CreatedAt = a.CreatedAt
		ar.UpdatedAt = a.UpdatedAt
		r.Posts = append(r.Posts, ar)
	}
	r.PostsCount = count
	return r
}

type commentResponse struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AuthorID  uint      `json:"author_id"`
	PostID    uint      `json:"post_id"`
}

type singleCommentResponse struct {
	Comment *commentResponse `json:"comment"`
}

type commentListResponse struct {
	Comments      []commentResponse `json:"comments"`
	CommentsCount int               `json:"commentsCount"`
}

func newCommentResponse(c echo.Context, cm *model.Comment) *singleCommentResponse {
	comment := new(commentResponse)
	comment.ID = cm.ID
	comment.AuthorID = cm.AuthorID
	comment.PostID = cm.PostID
	comment.Body = cm.Body
	comment.CreatedAt = cm.CreatedAt
	comment.UpdatedAt = cm.UpdatedAt
	return &singleCommentResponse{comment}
}

func newCommentListResponse(c echo.Context, comments []model.Comment, count int) *commentListResponse {
	r := new(commentListResponse)
	cr := commentResponse{}
	r.Comments = make([]commentResponse, 0)
	for _, i := range comments {
		cr.ID = i.ID
		cr.Body = i.Body
		cr.AuthorID = i.AuthorID
		cr.PostID = i.PostID
		cr.CreatedAt = i.CreatedAt
		cr.UpdatedAt = i.UpdatedAt

		r.Comments = append(r.Comments, cr)
	}
	r.CommentsCount = count
	return r
}
