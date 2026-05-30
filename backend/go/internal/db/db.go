package db

import (
	"log"

	"watch/backend/internal/config"
	"watch/backend/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg config.Config) *gorm.DB {
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}
	database, err := gorm.Open(mysql.Open(cfg.DBDSN), gormCfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("failed to get generic db: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	if err := database.AutoMigrate(
		&models.Admin{},
		&models.Role{},
		&models.Brand{},
		&models.Product{},
		&models.Banner{},
		&models.ContentPage{},
		&models.Order{},
	); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return database
}
