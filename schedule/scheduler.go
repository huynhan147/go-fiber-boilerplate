package schedule

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

// Register đăng ký tất cả scheduled tasks vào scheduler
func Register(s gocron.Scheduler) {
	// Chạy mỗi ngày lúc 00:00 — dọn dẹp token hết hạn
	s.NewJob(
		gocron.CronJob("0 0 * * *", false),
		gocron.NewTask(cleanExpiredTokens),
	)

	// Chạy mỗi giờ — gửi digest email
	s.NewJob(
		gocron.CronJob("0 * * * *", false),
		gocron.NewTask(sendHourlyDigest),
	)

	// Chạy mỗi phút — health check nội bộ
	s.NewJob(
		gocron.CronJob("* * * * *", false),
		gocron.NewTask(healthPing),
	)

	log.Println("📅 Scheduler registered")
}

func cleanExpiredTokens() {
	log.Println("🧹 [schedule] Cleaning expired tokens...")
	// TODO: xóa record token hết hạn trong DB
}

func sendHourlyDigest() {
	log.Println("📨 [schedule] Sending hourly digest...")
	// TODO: query users cần nhận digest, enqueue job gửi mail
}

func healthPing() {
	// Giữ im lặng để không spam log — chỉ log khi lỗi
	// log.Println("💓 [schedule] Health ping")
}
