package db

import (
	"gwentgg/services"

	"gorm.io/gorm"
)

type FactionGameStats struct {
	gorm.Model
	UserID      string
	Faction     services.Faction
	WinsCount   uint
	LossesCount uint
	DrawsCount  uint
	GamesCount  uint
}
