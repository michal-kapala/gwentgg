package models

import (
	"regexp"
)

type Deck struct {
	PlayerID      string `gorm:"primaryKey"`
	GameID        string `gorm:"primaryKey"`
	DeckID        uint
	ProvisionCost uint
	Leader        uint
	AbilityID     uint
	Content       string
	Vanities      string
	DeckHash      string
	DeckUntrusted bool
}

func (deck Deck) Parse() []string {
	// returns template IDs
	// leader, stratagem, other cards (unsorted)
	regex := regexp.MustCompile(`\b[0-9]{6}\b`)
	return regex.FindAllString(deck.Content, -1)
}

type DeckView struct {
	Leader    CardDefinition
	Stratagem CardDefinition
	Deck      []DeckCard
}

type DeckCard struct {
	Card   CardDefinition
	Copies uint
}
