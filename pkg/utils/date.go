package utils

import "time"

var Now = func() time.Time { return time.Now().UTC() }

func GetCurrentDate() string {
	now := Now()
	formatedDate := now.Format("2006-01-02 15:04:05")

	return formatedDate
}

func GetStringFromDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
