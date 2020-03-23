package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow - returns the current UTC time object
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString - returns the current UTC time as YYYY-MM-DD HH:MM:SS
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
