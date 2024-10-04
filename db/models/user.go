package models

import (
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
