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
	tagQuery := database.NewTagQuery(db, rdb)
	authorQuery := database.NewAuthorQuery(db, rdb)

	postService := service.NewPostService(postQuery)
	tagService := service.NewTagService(tagQuery)
	authorService := service.NewAuthorService(authorQuery)

	postHandler := handler.NewPostHandler(postService)
	tagHandler := handler.NewTagHandler(tagService)
	authorHandler := handler.NewAuthorHandler(authorService)

	// 注册路由
	apiv1 := r.Group("/api/v1")
	author := apiv1.Group("/author")
	{
		author.GET("/:author_id", authorHandler.GetAuthorDataById)
	}
	post := apiv1.Group("/post")
	{
		post.GET("", postHandler.GetPostData)
		post.GET("/recent", postHandler.GetRecentPostData)
		post.GET("/hot", postHandler.GetHotPostData)
		post.GET("/:post_id", postHandler.GetPostDataById)
	}
	tag := apiv1.Group("/tag")
	{
		tag.GET("", tagHandler.GetTagList)
	}
	return r
}
