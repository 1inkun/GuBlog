package utils

import (
	"net/http"

	"github.com/1inkun/GuBlog/internal/types"
	"github.com/labstack/echo/v5"
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

type Pagination struct {
	Data     any   `json:"data"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type BusinessErrResp struct {
	ErrCode int    `json:"errCode"`
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Detail  string `json:"detail"`
}

func NewResponse(c *echo.Context, msg string, data any) error {
	res := Response{
		Code:   0,
		Status: "success",
		Msg:    msg,
		Data:   data,
	}
	return c.JSON(http.StatusOK, res)
}

func NewBusinessResp(c *echo.Context, e types.BusinessError) error {
	res := BusinessErrResp{
		ErrCode: int(e.ErrCode),
		Status:  "faile",
		Msg:     e.Msg,
		Detail:  e.Detail,
	}
	return c.JSON(e.HttpCode, res)
}

func NewInternallErrResp(c *echo.Context) error {
	return c.JSON(http.StatusInternalServerError, "Internal Server Error")
}
