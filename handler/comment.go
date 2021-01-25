package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/model"
	"github.com/sOM2H/golang_trainee_backend/utils"
)

func (h *Handler) GetComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	a, err := h.commentStore.GetCommentByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newCommentResponse(c, a))
}

func (h *Handler) CreateComment(c echo.Context) error {
	var m model.Comment

	req := &createCommentRequest{}
	fmt.Println("Create Comment")
	fmt.Println(req.Comment.PostID)
	fmt.Println(req.Comment.Body)
	if err := req.bind(c, &m); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	fmt.Println("After Bind")
	fmt.Println(req.Comment.PostID)
	fmt.Println(req.Comment.Body)

	m.AuthorID = userIDFromToken(c)

	err := h.commentStore.CreateComment(&m)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newCommentResponse(c, &m))
}

func (h *Handler) UpdateComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	m, err := h.commentStore.GetCommentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if m == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req := &commentUpdateRequest{}

	if err := req.bind(c, m); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newCommentResponse(c, m))
}

func (h *Handler) GetComments(c echo.Context) error {
	var (
		comments []model.Comment
	)

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		offset = 0
	}

	comments, count, err := h.commentStore.ListComment(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, newCommentListResponse(c, comments, count))

}

// DeleteComment godoc
// @Summary Delete a comment for an post
// @Description Delete a comment for an post. Auth is required
// @ID delete-comments
// @Tags comment
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug of the post that you want to delete a comment for"
// @Param id path integer true "ID of the comment you want to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /posts/{slug}/comments/{id} [delete]
func (h *Handler) DeleteComment(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	cm, err := h.commentStore.GetCommentByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if cm == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	if cm.AuthorID != userIDFromToken(c) {
		return c.JSON(http.StatusUnauthorized, utils.NewError(errors.New("unauthorized action")))
	}

	if err := h.commentStore.DeleteComment(cm); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
