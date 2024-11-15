package handlers

import (
	"fmt"
	"gwentgg/components/pages"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/stats"

	"github.com/gofiber/fiber/v3"
	"strings"
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
	playerDeck := models.DeckView{
		Deck: []models.DeckCard{},
	}
	database.First(&playerDeck.Leader, "template_id = ?", ids[0])
	database.First(&playerDeck.Stratagem, "template_id = ?", ids[1])
	cards := []models.CardDefinition{}
	database.Order("provision_cost DESC, power DESC").Find(&cards, "template_id IN ? AND premium=0", ids[2:])
	for _, card := range cards {
		deckCard := models.DeckCard{
			Card:   card,
			Copies: uint(strings.Count(player.Deck.Content, card.TemplateID)),
		}
		playerDeck.Deck = append(playerDeck.Deck, deckCard)
	}
	ids = opponent.Deck.Parse()
	opponentDeck := models.DeckView{}
	database.First(&opponentDeck.Leader, "template_id = ?", ids[0])
	database.First(&opponentDeck.Stratagem, "template_id = ?", ids[1])
	cards = []models.CardDefinition{}
	database.Order("provision_cost DESC, power DESC").Find(&cards, "template_id IN ? AND premium=0", ids[2:])
	for _, card := range cards {
		deckCard := models.DeckCard{
			Card:   card,
			Copies: uint(strings.Count(opponent.Deck.Content, card.TemplateID)),
		}
		opponentDeck.Deck = append(opponentDeck.Deck, deckCard)
	}

	if player.PlayerID != playerID || game.ID != gameID {
		return c.SendStatus(404)
	}

	playerUser := models.User{}
	database.First(&playerUser, "id = ?", playerID)
	token := c.Cookies("access_token", "")
	season := c.Cookies("current_season", "")
	oppInfoResp, err := stats.Get(config.Get(c), opponent.PlayerID, token, season)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Opponent info request error, see server log for details.")
	}

	if token == "" || oppInfoResp.Status() != "200 OK" {
		return c.Redirect().To("/login")
	}

	opponentStats := oppInfoResp.Result().(*stats.RankedStats)
	opponentUser, err := opponentStats.ToModel()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save opponent data.")
	}
	db.ResetPlayerStats(database, opponent.PlayerID)
	database.Save(&opponentUser)

	return Render(c, pages.Game(&game, &player, &opponent, &playerDeck, &opponentDeck, &playerUser, &opponentUser))
}
