package games

import (
	"fmt"
	"gwentgg/db/models"
	"gwentgg/enums"
	"gwentgg/services"
)

type GameList struct {
	TotalCount    uint   `json:"total_count"`
	Limit         uint   `json:"limit"`
	PageToken     uint   `json:"page_token"`
	NextPageToken *uint  `json:"next_page_token"`
	Items         []Game `json:"items"`
}

type Game struct {
	ID              string     `json:"id"`
	GameType        string     `json:"game_type"`
	DateStarted     string     `json:"date_started"`
	DateFinished    string     `json:"date_finished"`
	GameServer      GameServer `json:"game_server"`
	Finished        string     `json:"finished"`
	Winner          *string    `json:"winner"`
	GoodGameAllowed bool       `json:"good_game_allowed"`
	Valid           bool       `json:"is_valid"`
	Players         []Player   `json:"players"`
}

func (game Game) ToModel() (models.Game, error) {
	started, err := services.ParseDate(game.DateStarted)
	if err != nil {
		return models.Game{}, err
	}

	finished, err := services.ParseDate(game.DateFinished)
	if err != nil {
		return models.Game{}, err
	}

	players := []models.GamePlayer{}
	for _, player := range game.Players {
		svPlayer := game.GameServer.Players[0]
		id := fmt.Sprintf("%d", svPlayer.ID)
		var playerModel models.GamePlayer
		if player.ID == id {
			playerModel, _ = player.ToModel(game.ID, svPlayer)
		} else {
			svPlayer = game.GameServer.Players[1]
			playerModel, _ = player.ToModel(game.ID, svPlayer)
		}
		players = append(players, playerModel)
	}

	model := models.Game{
		ID:                 game.ID,
		Type:               enums.GameType(game.GameType),
		DateStarted:        started,
		DateFinished:       finished,
		Finished:           enums.GameEnd(game.Finished),
		Winner:             game.Winner,
		GoodGameAllowed:    game.GoodGameAllowed,
		Valid:              game.Valid,
		GameVersion:        game.GameServer.GameVersion,
		VotingPatchVersion: game.GameServer.VotingPatchVersion,
		DataVersion:        game.GameServer.DataVersion,
		LogicVersion:       game.GameServer.LogicVersion,
		PackageVersion:     game.GameServer.PackageVersion,
		Server:             game.GameServer.Hostname,
		ServerGameID:       game.GameServer.GameID,
		StartingPlayer:     game.GameServer.StartingPlayer,
		FinishedRounds:     game.GameServer.FinishedRounds,
		Players:            players,
	}
	return model, err
}

type GameServer struct {
	ID                 uint           `json:"id"`
	GameVersion        string         `json:"game_version"`
	VotingPatchVersion string         `json:"voting_patch_version"`
	Hostname           string         `json:"hostname"`
	PackageVersion     string         `json:"package_version"`
	DataVersion        string         `json:"data_version"`
	LogicVersion       string         `json:"logic_version"`
	Network            string         `json:"network"`
	GameID             uint           `json:"game_id"`
	StartingPlayer     *uint          `json:"starting_player"`
	Players            []ServerPlayer `json:"players"`
	DateStarted        string         `json:"date_started"`
	DateFinished       string         `json:"date_finished"`
	FinishedRounds     uint           `json:"finished_rounds"`
}

type ServerPlayer struct {
	ID                   uint         `json:"Id"`
	InGamePlayerID       string       `json:"InGamePlayerId"`
	Rank                 uint         `json:"Rank"`
	MMR                  uint         `json:"MMR"`
	PlayedCards          []PlayedCard `json:"PlayedCards"`
	CardActions          []CardAction `json:"CardActions"`
	TrackedActionFilters string       `json:"TrackedActionFilters"`
	Deck                 Deck         `json:"Deck"`
	Coins                Coins        `json:"Coins"`
	Status               string       `json:"status"`
	BuildRegion          string       `json:"BuildRegion"`
	SessionID            string       `json:"SessionId"`
}

type PlayedCard struct {
	ID uint `json:"Id"`
}

type CardAction struct {
	Action      string  `json:"Action"`
	Amount      uint    `json:"Amount"`
	SourceType  string  `json:"SourceType"`
	SourceID    uint    `json:"SourceId"`
	SourceOwner string  `json:"SourceOwner"`
	TargetID    *uint   `json:"TargetId"`
	TargetOwner *string `json:"TargetOwner"`
	RoundID     uint    `json:"RoundId"`
	TurnID      uint    `json:"TurnId"`
}

func (action CardAction) ToModel(gameID string, playerID string, index uint) (models.CardAction, error) {
	cardAction := models.CardAction{
		ID:          index,
		PlayerID:    playerID,
		GameID:      gameID,
		Action:      enums.CardAction(action.Action),
		Amount:      action.Amount,
		SourceType:  enums.CardActionSource(action.SourceType),
		SourceID:    action.SourceID,
		SourceOwner: action.SourceOwner,
		TargetID:    action.TargetID,
		TargetOwner: action.TargetOwner,
		RoundID:     action.RoundID,
		TurnID:      action.TurnID,
	}
	return cardAction, nil
}

type Deck struct {
	DeckID        uint   `json:"DeckId"`
	ProvisionCost uint   `json:"ProvisionCost"`
	Leader        Leader `json:"Leader"`
	AbilityID     uint   `json:"AbilityId"`
	Content       string `json:"Content"`
	Vanities      []uint `json:"Vanities"`
	DeckHash      string `json:"DeckHash"`
	DeckUntrusted bool   `json:"DeckUntrusted"`
}

func (deck Deck) ToModel() (models.Deck, error) {
	vanities := ""
	for _, vanity := range deck.Vanities {
		if len(vanities) > 0 {
			vanities += ","
		}
		vanities += fmt.Sprintf("%d", vanity)
	}

	model := models.Deck{
		DeckID:        deck.DeckID,
		ProvisionCost: deck.ProvisionCost,
		Leader:        deck.Leader.ID,
		AbilityID:     deck.AbilityID,
		Content:       deck.Content,
		Vanities:      vanities,
		DeckHash:      deck.DeckHash,
		DeckUntrusted: deck.DeckUntrusted,
	}
	return model, nil
}

type Coins struct {
	Balance uint `json:"Balance"`
	Spent   uint `json:"Spent"`
}

type Leader struct {
	ID uint `json:"Id"`
}

type Player struct {
	ID           string  `json:"id"`
	RoundScores  []uint  `json:"round_scores"`
	RoundsWon    []bool  `json:"rounds_won"`
	Status       string  `json:"status"`
	DeckFaction  *string `json:"deck_faction"`
	SentGoodGame bool    `json:"sent_good_game"`
	Platform     string  `json:"platform"`
}

func (player Player) ToModel(gameID string, svPlayer ServerPlayer) (models.GamePlayer, error) {
	playedCards := ""
	for _, card := range svPlayer.PlayedCards {
		if len(playedCards) > 0 {
			playedCards += ","
		}
		playedCards += fmt.Sprintf("%d", card.ID)
	}

	actions := []models.CardAction{}
	for idx, action := range svPlayer.CardActions {
		actionModel, _ := action.ToModel(gameID, player.ID, uint(idx+1))
		actions = append(actions, actionModel)
	}

	deck, _ := svPlayer.Deck.ToModel()

	scores := ""
	for _, score := range player.RoundScores {
		if len(scores) > 0 {
			scores += ","
		}
		scores += fmt.Sprintf("%d", score)
	}

	wins := ""
	for _, win := range player.RoundsWon {
		if len(wins) > 0 {
			wins += ","
		}
		wins += fmt.Sprintf("%v", win)
	}

	model := models.GamePlayer{
		PlayerID:             player.ID,
		GameID:               gameID,
		InGamePlayerID:       svPlayer.InGamePlayerID,
		Rank:                 svPlayer.Rank,
		MMR:                  svPlayer.MMR,
		PlayedCards:          playedCards,
		CardActions:          actions,
		TrackedActionFilters: svPlayer.TrackedActionFilters,
		Deck:                 deck,
		CoinsBalance:         svPlayer.Coins.Balance,
		CoinsSpent:           svPlayer.Coins.Spent,
		Status:               enums.PlayerStatus(player.Status),
		SessionID:            svPlayer.SessionID,
		RoundScores:          scores,
		RoundsWon:            wins,
		DeckFaction:          player.DeckFaction,
		SentGoodGame:         player.SentGoodGame,
		Platform:             player.Platform,
	}
	return model, nil
}
