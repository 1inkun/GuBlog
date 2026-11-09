package router

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/1inkun/GuBlog/internal/database"
	"github.com/1inkun/GuBlog/internal/handler"
	"github.com/1inkun/GuBlog/internal/middlewares"
	"github.com/1inkun/GuBlog/internal/service"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func InitRouter(db *sql.DB, rdb *redis.Client) *echo.Echo {
	r := echo.New()
	// 配置错误处理
	r.HTTPErrorHandler = middlewares.HTTPErrorHandler
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// 配置限流
	r.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10.0)))
	// 配置logger、recover中间件
	r.Use(middleware.Recover())
	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		HandleError: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.Error("request failed", "uri", v.URI, "status", v.Status, "error", v.Error)
				return nil
			}
			logger.Info("request", "uri", v.URI, "status", v.Status)
			return nil
		},
	}))

	// 探针
	r.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello Web!")
	})

	// 初始化控制器
	postQuery := database.NewPostQuery(db, rdb)
	postService := service.NewPostService(postQuery)
	postHandler := handler.NewPostHandler(postService)

	// 注册路由
	apiv1 := r.Group("/api/v1")
	post := apiv1.Group("/post")
	{
		post.GET("", postHandler.GetPostData)
		post.GET("/:post_id", postHandler.GetPostDataById)
	}
	return r
}
