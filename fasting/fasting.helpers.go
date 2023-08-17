package fasting

import "time"

func getStartOfDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
}

func addDays(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}
