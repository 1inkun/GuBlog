package middlewares

import (
	"github.com/1inkun/GuBlog/internal/types"
	"github.com/1inkun/GuBlog/internal/utils"
	"github.com/labstack/echo/v5"
)

func HTTPErrorHandler(c *echo.Context, err error) {
	// 检查错误是否已经发送
	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return // already sent by a handler/middleware
		}
	}
	// 错误是否为业务错误
	if be, ok := err.(*types.BusinessError); ok {
		// 处理业务错误
		// log.Printf("%s", be)
		utils.NewBusinessResp(c, *be)
		return
	} else {
		// 处理非业务错误
		// log.Printf("内部错误:%s", err)
		utils.NewInternallErrResp(c)
		return
	}
}
