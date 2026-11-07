package handler

import (
	"log"
	"net/http"

	"github.com/1inkun/GuBlog/internal/service"
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

func (h *PostHandler) GetPostLists(c *echo.Context) error {
	log.Println("日志中间件测试信息")
	return c.String(http.StatusOK, "test")
}
