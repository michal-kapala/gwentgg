package models

import (
	"time"

	"gwentgg/enums"

	"gorm.io/gorm"
)

type CardDefinition struct {
	gorm.Model
	ID                string `gorm:"primaryKey"`
	TemplateID        string
	ArtID             string
	Name              string
	Faction           enums.Faction
	Power             uint
	CardGroup         enums.CardGroup
	ProvisionCost     uint
	Type              enums.CardType
	PrimaryCategoryID uint
	CategoryIDs       string
	SecondaryFactions string
	Armor             uint
	Premium           bool
	Availability      enums.CardSet
	Rarity            enums.CardRarity
	DateCreated       time.Time
	IsDeleted         bool
}
