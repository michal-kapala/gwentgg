package handlers

import (
	"fmt"
	"gwentgg/components/pages"
	"gwentgg/config"
	"gwentgg/db"
	"gwentgg/db/models"
	"gwentgg/services/stats"
	"gwentgg/utils"

	"sort"
	"strings"

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
	// decks
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
	// opponent stats
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
	// cards played
	playerCardIDs := strings.Split(player.PlayedCards, ",")
	playerCards := []models.CardDefinition{}
	for _, id := range playerCardIDs {
		var card models.CardDefinition
		database.First(&card, id)
		playerCards = append(playerCards, card)
	}
	opponentCardIDs := strings.Split(opponent.PlayedCards, ",")
	opponentCards := []models.CardDefinition{}
	for _, id := range opponentCardIDs {
		var card models.CardDefinition
		database.First(&card, id)
		opponentCards = append(opponentCards, card)
	}

	// card actions
	defIDs := []string{}
	for _, action := range player.CardActions {
		defIDs = append(defIDs, utils.ToString(action.SourceID))
		if action.TargetID != nil {
			defIDs = append(defIDs, utils.ToString(*action.TargetID))
		}
	}
	for _, action := range opponent.CardActions {
		defIDs = append(defIDs, utils.ToString(action.SourceID))
		if action.TargetID != nil {
			defIDs = append(defIDs, utils.ToString(*action.TargetID))
		}
	}

	defs := []models.CardDefinition{}
	database.Find(&defs, "id IN ?", defIDs)

	playerActions := []models.CardActionView{}
	for idx := range player.CardActions {
		view := models.CardActionView{
			Action: &player.CardActions[idx],
			Source: player.CardActions[idx].FindSource(defs),
			Target: player.CardActions[idx].FindTarget(defs),
			TurnOf: playerUser.Username,
		}
		playerActions = append(playerActions, view)
	}

	opponentActions := []models.CardActionView{}
	for idx := range opponent.CardActions {
		view := models.CardActionView{
			Action: &opponent.CardActions[idx],
			Source: opponent.CardActions[idx].FindSource(defs),
			Target: opponent.CardActions[idx].FindTarget(defs),
			TurnOf: opponentUser.Username,
		}
		opponentActions = append(opponentActions, view)
	}

	actions := append(playerActions, opponentActions...)
	sort.Slice(actions, func(i, j int) bool {
		if actions[i].Action.RoundID == actions[j].Action.RoundID {
			if actions[i].Action.TurnID == actions[j].Action.TurnID {
				return actions[i].Action.ID < actions[j].Action.ID
			}
			return actions[i].Action.TurnID < actions[j].Action.TurnID
		}
		return actions[i].Action.RoundID < actions[j].Action.RoundID
	})

	gameActions := models.GameCardActions{
		Round1: []models.CardActionView{},
		Round2: []models.CardActionView{},
		Round3: []models.CardActionView{},
	}
	for _, action := range actions {
		switch action.Action.RoundID {
		case 1:
			gameActions.Round1 = append(gameActions.Round1, action)
		case 2:
			gameActions.Round2 = append(gameActions.Round2, action)
		case 3:
			gameActions.Round3 = append(gameActions.Round3, action)
		}
	}

	return Render(c, pages.Game(&game, &player, &opponent, &playerDeck, &opponentDeck, &playerUser, &opponentUser, playerCards, opponentCards, &gameActions))
}
