package handler

import (
	"github.com/1inkun/GuBlog/internal/service"
	"github.com/1inkun/GuBlog/internal/types"
	"github.com/1inkun/GuBlog/internal/utils"
	"github.com/labstack/echo/v5"
)

type AuthorHandler struct {
	authorService *service.AuthorService
}

func NewAuthorHandler(authorService *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

func (h *AuthorHandler) GetAuthorDataById(c *echo.Context) error {
	ctx := c.Request().Context()
	authorId, err := echo.PathParam[int64](c, "author_id")
	if err != nil {
		return types.ErrBindingDataError
	}
	data, err := h.authorService.GetAuthorDataById(ctx, authorId)
	if err != nil {
		return err
	}
	return utils.NewResponse(c, "获取成功", data)
}
