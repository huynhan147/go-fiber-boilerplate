package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis(cfg *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetString("REDIS_HOST"), cfg.GetString("REDIS_PORT")),
		Password: cfg.GetString("REDIS_PASSWORD"),
		DB:       cfg.GetInt("REDIS_DB"),
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected")
	return rdb
}
