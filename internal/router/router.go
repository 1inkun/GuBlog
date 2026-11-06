package router

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v5"
)

func InitRouter(db *sql.DB) *echo.Echo {
	r := echo.New()
	r.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello Web!")
	})
	return r
}
