package db

import (
	"gwentgg/db/models"

	"github.com/gofiber/fiber/v3"
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
		&models.User{},
		&models.FactionGameStats{},
		&models.FactionProgression{},
		&models.Season{},
		&models.CardDefinition{},
		&models.Game{},
		&models.GamePlayer{},
		&models.CardAction{},
		&models.Deck{},
	)
	return err
}

func Add(dbCtx *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Locals("db", dbCtx)
		return c.Next()
	}
}

func Get(c fiber.Ctx) *gorm.DB {
	return c.Locals("db").(*gorm.DB)
}

func ResetPlayerStats(dbCtx *gorm.DB, userID string) {
	dbCtx.Where(&models.FactionGameStats{UserID: userID}).Unscoped().Delete(&models.FactionGameStats{})
	dbCtx.Where(&models.FactionProgression{UserID: userID}).Unscoped().Delete(&models.FactionProgression{})
}

func GetSeasonName(dbCtx *gorm.DB, seasonID string) string {
	var season models.Season
	dbCtx.Where(&models.Season{ID: seasonID}).First(&season)
	return season.Name
}
