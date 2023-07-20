package db

import (
	"github.com/sarahrajabazdeh/DreamPilot/model"
	"gorm.io/gorm"
)

// MigrateDatabase performs the database migration for all models
func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Goal{}); err != nil {
		return err
	}
	return nil
}
