package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.GetString("DB_USERNAME"),
		cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_HOST"),
		cfg.GetString("DB_PORT"),
		cfg.GetString("DB_DATABASE"),
		cfg.GetString("DB_CHARSET"),
	)

	gormCfg := &gorm.Config{}
	if cfg.GetString("APP_ENV") == "local" {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL: %v", err)
	}

	// Connection pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("✅ MySQL connected")
	return db
}
