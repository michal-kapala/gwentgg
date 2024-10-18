package utils

import (
	"fmt"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02T15:04:05-0700"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ToString(number uint) string {
	return fmt.Sprintf("%d", number)
}
