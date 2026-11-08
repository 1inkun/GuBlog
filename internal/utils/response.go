package utils

import (
	"net/http"

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

func NewResponse(c *echo.Context, msg string, data any) error {
	res := Response{
		Code:   0,
		Status: "success",
		Msg:    msg,
		Data:   data,
	}
	return c.JSON(http.StatusOK, res)
}
