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

// func (h *PostHandler) GetPostList(c *echo.Context) error {
// 	ctx := c.Request().Context()
// 	var req types.Pagination
// 	if err := c.Bind(&req); err != nil {
// 		return types.ErrBindingDataError
// 	}
// 	data, err := h.postService.GetPostList(ctx, req.Page, req.PageSize)
// 	if err != nil {
// 		return err
// 	}
// 	return utils.NewResponse(c, "获取成功", data)
// }

func (h *PostHandler) GetPostDataById(c *echo.Context) error {
	ctx := c.Request().Context()
	postId, err := echo.PathParam[int64](c, "post_id")
	if err != nil {
		return types.ErrBindingDataError
	}
	data, err := h.postService.GetPostDataById(ctx, postId)
	if err != nil {
		return err
	}
	return utils.NewResponse(c, "获取成功", data)
}

func (h *PostHandler) GetPostData(c *echo.Context) error {
	ctx := c.Request().Context()
	var req types.GetPostDataReq
	if err := c.Bind(&req); err != nil {
		return types.ErrBindingDataError
	}
	var data utils.Pagination
	if len(req.Tags) == 0 {
		res, err := h.postService.GetPostList(ctx, req.Pagination.Page, req.Pagination.PageSize)
		if err != nil {
			return err
		}
		data = res
	} else {
		res, err := h.postService.GetPostListWithTags(ctx, req.Tags, req.Pagination.Page, req.Pagination.PageSize)
		if err != nil {
			return err
		}
		data = res
	}
	return utils.NewResponse(c, "查询成功", data)
}
