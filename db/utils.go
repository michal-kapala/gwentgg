package db

import (
	"gwentgg/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(filename string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}

func AutoMigrate(database *gorm.DB) error {
	err := database.AutoMigrate(
		&db.User{},
		&db.FactionGameStats{},
		&db.FactionProgression{},
	)
	return err
}
