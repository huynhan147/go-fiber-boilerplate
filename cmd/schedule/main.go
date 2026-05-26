package main

import (
	"log"
	"myapp/config"
	"myapp/schedule"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	config.Load()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("❌ Failed to create scheduler: %v", err)
	}

	// Đăng ký tất cả scheduled tasks
	schedule.Register(s)

	// Start non-blocking
	s.Start()
	log.Println("📅 Scheduler started")

	// Graceful shutdown khi nhận SIGINT / SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Scheduler shutting down...")
	if err := s.Shutdown(); err != nil {
		log.Fatalf("❌ Scheduler shutdown error: %v", err)
	}
}
