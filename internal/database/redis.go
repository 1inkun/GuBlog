package database

import (
	"context"
	"log"

	"github.com/1inkun/GuBlog/config"
	"github.com/redis/go-redis/v9"
)

func InitRedisConn(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Hostname + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis连接失败:%s", err)
		return nil, err
	}
	return rdb, nil
}
