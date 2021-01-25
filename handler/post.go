package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/model"
	"github.com/sOM2H/golang_trainee_backend/utils"
)

// GetPost godoc
// @Summary Get an post
// @Description Get an post. Auth not required
// @ID get-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the post to get"
// @Success 200 {object} singlePostResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /posts/{slug} [get]
func (h *Handler) GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	a, err := h.postStore.GetPostByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newPostResponse(c, a))
}

// Posts godoc
// @Summary Get recent posts globally
// @Description Get most recent posts globally. Use query parameters to filter results. Auth is optional
// @ID get-posts
// @Tags post
// @Accept  json
// @Produce  json
// @Param tag query string false "Filter by tag"
// @Param author query string false "Filter by author (username)"
// @Param favorited query string false "Filter by favorites of a user (username)"
// @Param limit query integer false "Limit number of posts returned (default is 20)"
// @Param offset query integer false "Offset/skip number of posts (default is 0)"
// @Success 200 {object} postListResponse
// @Failure 500 {object} utils.Error
// @Router /posts [get]
func (h *Handler) GetPosts(c echo.Context) error {
	var (
		posts []model.Post
		count int
	)

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	posts, count, err = h.postStore.ListPost(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, newPostListResponse(h.userStore, userIDFromToken(c), posts, count))
}

// CreatePost godoc
// @Summary Create an post
// @Description Create an post. Auth is require
// @ID create-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param post body postCreateRequest true "Post to create"
// @Success 201 {object} singlePostResponse
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /posts [post]
func (h *Handler) CreatePost(c echo.Context) error {
	var a model.Post

	req := &postCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	a.AuthorID = userIDFromToken(c)

	err := h.postStore.CreatePost(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newPostResponse(c, &a))
}

// UpdatePost godoc
// @Summary Update an post
// @Description Update an post. Auth is required
// @ID update-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the post to update"
// @Param post body postUpdateRequest true "Post to update"
// @Success 200 {object} singlePostResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /posts/{slug} [put]
func (h *Handler) UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	a, err := h.postStore.GetPostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req := &postUpdateRequest{}
	req.populate(a)

	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newPostResponse(c, a))
}

// DeletePost godoc
// @Summary Delete an post
// @Description Delete an post. Auth is required
// @ID delete-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the post to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /posts/{slug} [delete]
func (h *Handler) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	a, err := h.postStore.GetPostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	err = h.postStore.DeletePost(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
