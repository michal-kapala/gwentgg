package models

import (
	"fmt"
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

func (player GamePlayer) MakeLeaderAssetPath() string {
	return fmt.Sprintf("/assets/leaders/%d.png", player.Deck.Leader)
}

func (player GamePlayer) MakeLeaderBgAsset() string {
	switch enums.Faction(*player.DeckFaction) {
	case enums.Monsters:
		return "mon-background-right.png"
	case enums.Nilfgaard:
		return "nilf-background-right.png"
	case enums.NorthernKingdoms:
		return "nor-background-right.png"
	case enums.Scoiatael:
		return "sco-background-right.png"
	case enums.Skellige:
		return "ske-background-right.png"
	case enums.Syndicate:
		return "syn-background-right.png"
	}
	return ""
}
