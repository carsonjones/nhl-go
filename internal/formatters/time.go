package formatters

import (
	"fmt"
	"time"
)

// FormatGameTime formats a UTC time string into EST and CST times
func FormatGameTime(utcTime string) (string, error) {
	// Parse the UTC time
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}

	// Load EST location
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		return "", fmt.Errorf("error loading EST location: %v", err)
	}

	// Load CST location
	cst, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return "", fmt.Errorf("error loading CST location: %v", err)
	}

	// Convert to EST and CST
	estTime := t.In(est)
	cstTime := t.In(cst)

	// Format the times
	return fmt.Sprintf("%s EST (%s CST)",
		estTime.Format("3:04 PM"),
		cstTime.Format("3:04 PM")), nil
}

// FormatTimeOnIce formats seconds into a time on ice string (MM:SS)
func FormatTimeOnIce(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, remainingSeconds)
}
