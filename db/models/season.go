package models

import (
	"time"

	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Name         string
	DateStarts   time.Time
	DateEnds     time.Time
	DateCreated  time.Time
	RankingType  string
	ParentSeason *string
	SeasonType   *string
}
