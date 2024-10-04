package models

import (
	"gwentgg/enums"

	"gorm.io/gorm"
)

type FactionProgression struct {
	gorm.Model
	ID                               int `gorm:"primaryKey"`
	UserID                           string
	Faction                          enums.Faction
	IsUsedForScoreCalculation        bool
	UnlockedScore                    int
	RealScore                        int
	RealScorePeak                    int
	GamesCount                       uint
	UnlockedScoreGamesCountThreshold uint
	UnlockedScoreFraction            float64
}
