package models

import (
	"fmt"
	"gwentgg/enums"
	"math"

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

func (stats FactionGameStats) Winrate(precision uint) float64 {
	if stats.GamesCount == 0 {
		return -1.0
	}
	factor := math.Pow(10.0, float64(precision))
	wr := factor * 100 * float64(stats.WinsCount) / float64(stats.GamesCount)
	wr = math.Round(wr)
	wr /= factor
	return wr
}

func (stats FactionGameStats) WinrateStr(precision uint) string {
	wr := stats.Winrate(precision)
	if wr == -1.0 {
		return "N/A"
	}
	return fmt.Sprintf("%.1f", wr)
}
