package models

import (
	"gwentgg/enums"
)

type GamePlayer struct {
	PlayerID             string `gorm:"primaryKey"`
	GameID               string `gorm:"primaryKey"`
	InGamePlayerID       string
	Rank                 uint
	MMR                  uint
	PlayedCards          string
	CardActions          []CardAction `gorm:"foreignKey:GameID,PlayerID"`
	TrackedActionFilters string
	Deck                 Deck `gorm:"foreignKey:GameID,PlayerID"`
	CoinsBalance         uint
	CoinsSpent           uint
	Status               enums.PlayerStatus
	SessionID            string
	RoundScores          string
	RoundsWon            string
	DeckFaction          *string
	SentGoodGame         bool
	Platform             string
}
