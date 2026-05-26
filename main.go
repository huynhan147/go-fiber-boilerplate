package main

import (
	"log"
	"myapp/app/bootstrap"
	"myapp/app/middleware"
	"myapp/config"
	"myapp/pkg/logger"
	"myapp/routes"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Load config
	cfg := config.Load()

	// Init logger — phải gọi trước tất cả
	logger.Init(
		cfg.GetString("APP_ENV"),
		cfg.GetString("LOG_DIR"),
	)
	defer logger.Sync()

	logger.Info("Starting application",
		logger.F("app", cfg.GetString("APP_NAME")),
		logger.F("env", cfg.GetString("APP_ENV")),
	)

	// Init database
	db := config.InitDB(cfg)

	// Init redis
	rdb := config.InitRedis(cfg)

	// Build container
	container := bootstrap.BuildContainer(
		db,
		rdb,
		cfg,
	)

	// Init Fiber
	app := fiber.New(fiber.Config{
		AppName:      cfg.GetString("APP_NAME"),
		ErrorHandler: middleware.ErrorHandler,
	})

	// Global middleware
	app.Use(recover.New())

	app.Use(fiberlogger.New(fiberlogger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     config.AccessLogWriter(cfg.GetString("LOG_DIR")),
	}))

	app.Use(fiberlogger.New(fiberlogger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	app.Use(middleware.CORS())

	// Register routes
	routes.Register(app, container)

	// Start server
	port := cfg.GetString("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on http://localhost:%s", port)

	log.Fatal(app.Listen(":" + port))
}
