package src

import (
	"time"
)

func FormatDate(date time.Time) string {
	return date.Weekday().String() + " " + date.Format(time.DateOnly)
}

func GetDayFromDate(date time.Time) string {
	return date.Weekday().String()
}

func GetDateFromDate(date time.Time) string {
	return date.Format(time.DateOnly)
}
