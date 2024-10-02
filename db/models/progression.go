package db

import (
	"gwentgg/enums"

	"gorm.io/gorm"
)

type FactionProgression struct {
	gorm.Model
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
