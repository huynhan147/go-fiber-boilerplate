package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func createMigration(name string) {
	if name == "" {
		log.Fatal("❌ Missing migration name. Usage: make migrate-create name=create_posts_table")
	}

	dir := "database/migrations"
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("❌ Cannot create migrations dir: %v", err)
	}

	// Đếm file hiện có để lấy số thứ tự tiếp theo
	entries, _ := filepath.Glob(filepath.Join(dir, "*.up.sql"))
	seq := len(entries) + 1

	prefix := fmt.Sprintf("%06d_%s", seq, name)
	upFile   := filepath.Join(dir, prefix+".up.sql")
	downFile := filepath.Join(dir, prefix+".down.sql")

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	upContent := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n-- TODO: write your UP migration here\n\n", prefix, timestamp)
	downContent := fmt.Sprintf("-- Migration: %s\n-- Created at: %s\n-- TODO: write your DOWN migration here\n\n", prefix, timestamp)

	if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
		log.Fatalf("❌ Failed to create up file: %v", err)
	}
	if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
		log.Fatalf("❌ Failed to create down file: %v", err)
	}

	fmt.Printf("✅ Created migration files:\n")
	fmt.Printf("   ↑  %s\n", upFile)
	fmt.Printf("   ↓  %s\n", downFile)
}
