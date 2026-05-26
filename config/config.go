package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Println("⚠️  No .env file found, reading from environment variables")
	}

	return v
}
