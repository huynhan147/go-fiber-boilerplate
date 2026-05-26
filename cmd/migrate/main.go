package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cmd := "up"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	// Lệnh create không cần connect DB
	if cmd == "create" {
		name := ""
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		createMigration(name)
		return
	}

	// Các lệnh còn lại cần connect DB
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&multiStatements=true",
		cfg.GetString("DB_USERNAME"),
		cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_PORT"),
		cfg.GetString("DB_DATABASE"),
		cfg.GetString("DB_CHARSET"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "mysql", driver)
	if err != nil {
		log.Fatalf("❌ Failed to init migrate: %v", err)
	}

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("❌ Migrate up failed: %v", err)
		}
		log.Println("✅ Migrate up done")

	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("❌ Migrate down failed: %v", err)
		}
		log.Println("✅ Rolled back 1 step")

	case "reset":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("❌ Migrate reset failed: %v", err)
		}
		log.Println("✅ All migrations rolled back")

	case "version":
		v, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("❌ Failed to get version: %v", err)
		}
		log.Printf("📌 Current version: %d, dirty: %v", v, dirty)

	default:
		log.Fatalf("❌ Unknown command: %s\nUsage: up | down | reset | version | create <name>", cmd)
	}
}
