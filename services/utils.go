package services

import (
	"time"

	"gorm.io/gorm"
)

type Model interface {
	ToModel() (gorm.Model, error)
}

// Model with a foreign key.
type Submodel interface {
	ToModel(fkey string) (gorm.Model, error)
}

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02T15:04:05-0700"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
