package handler

import (
	"github.com/1inkun/GuBlog/internal/service"
	"github.com/1inkun/GuBlog/internal/types"
	"github.com/1inkun/GuBlog/internal/utils"
	"github.com/labstack/echo/v5"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler(tagService *service.TagService) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

func (h *TagHandler) GetTagList(c *echo.Context) error {
	ctx := c.Request().Context()
	var req types.Pagination
	if err := c.Bind(&req); err != nil {
		return types.ErrBindingDataError
	}
	data, err := h.tagService.GetTagList(ctx, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return utils.NewResponse(c, "获取成功", data)
}
