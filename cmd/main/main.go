package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1inkun/GuBlog/config"
	"github.com/1inkun/GuBlog/internal/database"
	"github.com/1inkun/GuBlog/internal/router"
)

func main() {
	// 初始化配置
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("初始化配置出错:%s", err)
	}

	// 初始化数据库连接
	db, err := database.InitDatabaseConn()
	if err != nil {
		log.Fatalf("连接数据库出错:%s", err)
	}
	defer db.Close()
	rdb, err := database.InitRedisConn(cfg)
	if err != nil {
		log.Fatalf("连接Redis出错:%s", err)
	}
	defer rdb.Close()

	// 注册优雅关闭
	r := router.InitRouter(db, rdb)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 启动服务器
	go func() {
		log.Println("服务器启动成功，监听端口:8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("服务器启动失败:%s", err)
		}
	}()

	// 传递服务器关闭信号
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop server:%s", err)
	}
}
