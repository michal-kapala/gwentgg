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
	database.First(&game, "id = ?", gameID)
	players := []models.GamePlayer{}
	database.Preload("Deck").Preload("CardActions").Find(&players, "game_id = ?", gameID)
	game.Players = players
	if len(game.Players) != 2 {
		return c.SendStatus(404)
	}
	player := game.Players[0]
	opponent := game.Players[1]
	if player.PlayerID != playerID {
		player = game.Players[1]
		opponent = game.Players[0]
	}
	ids := player.Deck.Parse()
	playerDeck := models.DeckView{}
	database.First(&playerDeck.Leader, "template_id = ?", ids[0])
	database.First(&playerDeck.Stratagem, "template_id = ?", ids[1])
	database.Order("provision_cost DESC, power DESC").Find(&playerDeck.Deck, "template_id IN ?", ids[2:])
	ids = opponent.Deck.Parse()
	opponentDeck := models.DeckView{}
	database.First(&opponentDeck.Leader, "template_id = ?", ids[0])
	database.First(&opponentDeck.Stratagem, "template_id = ?", ids[1])
	database.Order("provision_cost DESC, power DESC").Find(&opponentDeck.Deck, "template_id IN ?", ids[2:])

	if player.PlayerID != playerID || game.ID != gameID {
		return c.SendStatus(404)
	}
	return Render(c, pages.Game(&game, &player, &playerDeck, &opponentDeck))
}
