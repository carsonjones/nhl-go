package formatters

import (
	"fmt"
	"time"
)

// FormatSeasonID formats a season ID into a readable string (e.g., "2023-2024")
func FormatSeasonID(seasonID int) string {
	start := seasonID / 10000
	return fmt.Sprintf("%d-%d", start, start+1)
}

// GetCurrentSeasonID returns the current NHL season ID
func GetCurrentSeasonID() int {
	now := time.Now()
	year := now.Year()
	// NHL season typically starts in October
	if now.Month() < time.October {
		year-- // If we're before October, we're in the previous year's season
	}
	return year*10000 + (year + 1)
}
