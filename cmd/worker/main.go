package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/jobs"

	"github.com/hibiken/asynq"
)

func main() {
	cfg := config.Load()

	redisOpt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetString("REDIS_HOST"), cfg.GetString("REDIS_PORT")),
		Password: cfg.GetString("REDIS_PASSWORD"),
		DB:       cfg.GetInt("REDIS_DB"),
	}

	// Cấu hình worker: 10 concurrent tasks, ưu tiên queue critical > default > low
	srv := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	// Đăng ký handler cho từng task type
	mux := asynq.NewServeMux()
	mux.Handle(jobs.TypeSendWelcomeEmail, jobs.NewSendWelcomeEmailHandler())
	mux.Handle(jobs.TypeProcessAvatar, jobs.NewProcessAvatarHandler())

	log.Println("🚀 Worker started")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("❌ Worker failed: %v", err)
	}
}
