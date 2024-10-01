package db

import (
	"gwentgg/services"

	"gorm.io/gorm"
)

type FactionProgression struct {
	gorm.Model
	UserID                           string
	Faction                          services.Faction
	IsUsedForScoreCalculation        bool
	UnlockedScore                    int
	RealScore                        int
	RealScorePeak                    int
	GamesCount                       uint
	UnlockedScoreGamesCountThreshold uint
	UnlockedScoreFraction            float64
}
