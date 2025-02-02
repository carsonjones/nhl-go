package main

import (
	"testing"
	"time"
)

func TestFormatGameTime(t *testing.T) {
	tests := []struct {
		name    string
		utcTime string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid UTC time",
			utcTime: "2024-02-15T00:00:00Z",
			want:    "7:00 PM EST (6:00 PM CST)",
			wantErr: false,
		},
		{
			name:    "Invalid UTC time",
			utcTime: "invalid",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatGameTime(tt.utcTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatGameTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("formatGameTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSeasonID(t *testing.T) {
	tests := []struct {
		name     string
		seasonID int
		want     string
	}{
		{
			name:     "Current season",
			seasonID: 20232024,
			want:     "2023-2024",
		},
		{
			name:     "Past season",
			seasonID: 20222023,
			want:     "2022-2023",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatSeasonID(tt.seasonID); got != tt.want {
				t.Errorf("formatSeasonID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeOnIce(t *testing.T) {
	tests := []struct {
		name    string
		seconds int
		want    string
	}{
		{
			name:    "One hour",
			seconds: 3600,
			want:    "1:00",
		},
		{
			name:    "One hour thirty minutes",
			seconds: 5400,
			want:    "1:30",
		},
		{
			name:    "Zero",
			seconds: 0,
			want:    "0:00",
		},
		{
			name:    "Less than one minute",
			seconds: 45,
			want:    "0:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTimeOnIce(tt.seconds); got != tt.want {
				t.Errorf("formatTimeOnIce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentSeasonID(t *testing.T) {
	// Since this function depends on the current date, we need to test different scenarios
	now := time.Now()
	year := now.Year()

	// If we're before October, the season should be the previous year
	if now.Month() < time.October {
		year--
	}

	got := getCurrentSeasonID()
	want := year*10000 + (year + 1)

	if got != want {
		t.Errorf("getCurrentSeasonID() = %v, want %v", got, want)
	}
}
