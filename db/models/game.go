package models

import (
	"gwentgg/enums"
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID                 string `gorm:"primaryKey"`
	Type               enums.GameType
	DateStarted        time.Time
	DateFinished       time.Time
	Finished           enums.GameEnd
	Winner             *string
	GoodGameAllowed    bool
	Valid              bool
	GameVersion        string
	VotingPatchVersion string
	DataVersion        string
	LogicVersion       string
	PackageVersion     string
	Server             string
	ServerGameID       uint
	StartingPlayer     *uint
	FinishedRounds     uint
	Players            []GamePlayer `gorm:"foreignKey:GameID"`
}
