package handlers

import (
	"gwentgg/components/pages"
	"gwentgg/db"
	"gwentgg/db/models"

	"github.com/gofiber/fiber/v3"
)

func GameHandler(c fiber.Ctx) error {
	gameID := c.Params("game")
	var game models.Game
	database := db.Get(c)
	database.Preload("Players").First(&game, "id = ?", gameID)
	return Render(c, pages.Game(&game))
}
