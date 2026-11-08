package handler

import (
	"github.com/1inkun/GuBlog/internal/service"
	"github.com/1inkun/GuBlog/internal/types"
	"github.com/1inkun/GuBlog/internal/utils"
	"github.com/labstack/echo/v5"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) GetPostList(c *echo.Context) error {
	ctx := c.Request().Context()
	var req types.Pagination
	if err := c.Bind(&req); err != nil {
		return err
	}
	data, err := h.postService.GetPostList(ctx, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return utils.NewResponse(c, "获取成功", data)
}
