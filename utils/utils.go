package utils

import (
	"fmt"
	"math"
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

func RelativeTime(date time.Time) string {
	now := time.Now()
	diff := now.Sub(date)
	minutes := uint(math.Round(diff.Minutes()))
	fmtStr := "%dmin ago"
	if minutes < 60 {
		return fmt.Sprintf(fmtStr, minutes)
	}
	hours := uint(math.Round(diff.Hours()))
	if hours < 24 {
		fmtStr = "%dh ago"
		return fmt.Sprintf("%dh ago", hours)
	}
	days := uint(math.Round(float64(hours) / 24.0))
	if days < 30 {
		fmtStr = "%d days ago"
		if days == 1 {
			fmtStr = "%d day ago"
		}
		return fmt.Sprintf(fmtStr, days)
	}
	months := uint(math.Round(float64(days) / 30.0))
	if months < 12 {
		fmtStr = "%d months ago"
		if months == 1 {
			fmtStr = "%d month ago"
		}
		return fmt.Sprintf(fmtStr, days)
	}
	years := uint(math.Round(float64(months) / 12.0))
	fmtStr = "%d year ago"
	if years > 1 {
		fmtStr = "%d years ago"
	}
	return fmt.Sprintf(fmtStr, years)
}
