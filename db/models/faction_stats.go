package db

import (
	"gwentgg/enums"

	"gorm.io/gorm"
)

type FactionGameStats struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	UserID      string
	Faction     enums.Faction
	WinsCount   uint
	LossesCount uint
	DrawsCount  uint
	GamesCount  uint
}
