package models

import (
	"gwentgg/enums"

	"gorm.io/gorm"
)

type CardAction struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:false"`
	PlayerID    string `gorm:"primaryKey"`
	GameID      string `gorm:"primaryKey"`
	Action      enums.CardAction
	Amount      uint
	SourceType  enums.CardActionSource
	SourceID    uint
	SourceOwner string
	TargetID    *uint
	TargetOwner *string
	RoundID     uint
	TurnID      uint
}
