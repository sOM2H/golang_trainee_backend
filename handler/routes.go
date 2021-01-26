package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/router/middleware"
	"github.com/sOM2H/golang_trainee_backend/utils"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

	posts := v1.Group("/posts", jwtMiddleware)
	posts.POST("", h.CreatePost)
	posts.GET("", h.GetPosts)
	posts.GET("/:id", h.GetPost)
	posts.PUT("/:id", h.UpdatePost)
	posts.DELETE("/:id", h.DeletePost)

	comments := v1.Group("/comments", jwtMiddleware)
	comments.POST("", h.CreateComment)
	comments.GET("", h.GetComment)
	comments.GET("/:id", h.ListCommentByPostID)
	comments.PUT("/:id", h.UpdateComment)
	comments.DELETE("/:id", h.DeleteComment)

}
