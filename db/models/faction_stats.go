package db

import (
	"gwentgg/enums"

	"gorm.io/gorm"
)

type FactionGameStats struct {
	gorm.Model
	UserID      string
	Faction     enums.Faction
	WinsCount   uint
	LossesCount uint
	DrawsCount  uint
	GamesCount  uint
}
