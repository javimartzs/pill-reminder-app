package database

import (
	"github.com/javimartzs/pill-reminder-app-v0/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Medicamento{})
}
