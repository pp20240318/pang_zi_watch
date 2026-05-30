package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string
	ServerPort string
	DBDSN      string
	JWTSecret  string
	JWTExpiry  time.Duration
}

func Load() Config {
	v := viper.New()
	v.SetEnvPrefix("WATCH")
	v.AutomaticEnv()

	v.SetDefault("ENV", "dev")
	v.SetDefault("SERVER_PORT", "8080")
	v.SetDefault("DB_DSN", "")
	v.SetDefault("JWT_SECRET", "")
	v.SetDefault("JWT_EXPIRY_MINUTES", 120)

	cfg := Config{
		Env:        v.GetString("ENV"),
		ServerPort: v.GetString("SERVER_PORT"),
		DBDSN:      v.GetString("DB_DSN"),
		JWTSecret:  v.GetString("JWT_SECRET"),
		JWTExpiry:  time.Duration(v.GetInt("JWT_EXPIRY_MINUTES")) * time.Minute,
	}

	if cfg.Env == "prod" {
		if cfg.JWTSecret == "" {
			log.Fatal("JWT secret cannot be empty in prod (set WATCH_JWT_SECRET)")
		}
		if cfg.DBDSN == "" {
			log.Fatal("DB DSN cannot be empty in prod (set WATCH_DB_DSN)")
		}
	} else {
		if cfg.JWTSecret == "" {
			cfg.JWTSecret = "change-this-secret"
			log.Printf("[config] using default dev JWT secret; set WATCH_JWT_SECRET to override")
		}
		if cfg.DBDSN == "" {
			cfg.DBDSN = "watch:Aa123456@tcp(127.0.0.1:3306)/watch?parseTime=true&loc=Local&charset=utf8mb4"
			log.Printf("[config] using default dev DB DSN; set WATCH_DB_DSN to override")
		}
	}

	return cfg
}

func (c Config) Addr() string {
	return fmt.Sprintf(":%s", c.ServerPort)
}
