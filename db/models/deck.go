package models

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
