package utils

import (
	"fmt"
	"time"
)

// ConvertDate converts dd-mm-yyyy into time.Time
func ConvertDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, fmt.Errorf("date cannot be empty")
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf(
			"invalid date format: expected dd-mm-yyyy, got %s",
			date,
		)
	}

	return t, nil
}
