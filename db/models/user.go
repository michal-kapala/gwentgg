package models

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                  string `gorm:"primaryKey"`
	Username            string
	Platform            string
	Title               string
	Score               int
	Position            uint
	ContinentalPosition uint
	RankID              uint
	Leaderboard         string
	DateScoreUpdated    time.Time
	DateLevelUpdated    time.Time
	DateCreated         time.Time
	WinsCount           uint
	LossesCount         uint
	DrawsCount          uint
	GamesCount          uint
	Level               uint
	Continent           string
	Country             string
	MosaicPieceCount    uint
	Rank                uint
	RatingScore         int
	ParagonLevel        int
	AvatarID            string
	BorderID            string
	TitleID             string
	Progressions        []FactionProgression
	FactionStats        []FactionGameStats
}

func (user User) Winrate(precision uint) float64 {
	if user.GamesCount == 0 {
		return -1.0
	}
	factor := math.Pow(10.0, float64(precision))
	wr := factor * 100 * float64(user.WinsCount) / float64(user.GamesCount)
	wr = math.Round(wr)
	wr /= factor
	return wr
}

func (user User) WinrateStr(precision uint) string {
	wr := user.Winrate(precision)
	if wr == -1.0 {
		return "N/A"
	}
	return fmt.Sprintf("%.2f", wr)
}
