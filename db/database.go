package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/javimartzs/pill-reminder-app-v0/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.EnvConfig) *gorm.DB {

	uri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s",
		config.DBHost, config.DBUser, config.DBName, config.DBPass, config.DBPort)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database %v", err)
	}

	log.Info("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate: %v", err)
	}

	return db
}
