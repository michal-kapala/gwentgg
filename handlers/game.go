package handlers

import (
	"gwentgg/components/pages"
	"gwentgg/db"
	"gwentgg/db/models"

	"github.com/gofiber/fiber/v3"
)

func GameHandler(c fiber.Ctx) error {
	playerID := c.Params("user")
	gameID := c.Params("game")
	database := db.Get(c)
	var game models.Game
	database.Preload("Players").First(&game, "id = ?", gameID)
	if len(game.Players) != 2 {
		return c.SendStatus(404)
	}
	player := game.Players[0]
	if player.PlayerID != playerID {
		player = game.Players[1]
	}
	if player.PlayerID != playerID || game.ID != gameID {
		return c.SendStatus(404)
	}	
	return Render(c, pages.Game(&game, &player))
}
