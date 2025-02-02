package main

import (
	"go-nhl/internal/formatters"
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
			utcTime: "2024-02-01T00:00:00Z",
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
			got, err := formatters.FormatGameTime(tt.utcTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatGameTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatGameTime() = %v, want %v", got, tt.want)
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
			if got := formatters.FormatSeasonID(tt.seasonID); got != tt.want {
				t.Errorf("FormatSeasonID() = %v, want %v", got, tt.want)
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
			name:    "One minute",
			seconds: 60,
			want:    "1:00",
		},
		{
			name:    "One hour",
			seconds: 3600,
			want:    "60:00",
		},
		{
			name:    "Complex time",
			seconds: 3665,
			want:    "61:05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatters.FormatTimeOnIce(tt.seconds); got != tt.want {
				t.Errorf("FormatTimeOnIce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentSeasonID(t *testing.T) {
	now := time.Now()
	year := now.Year()
	if now.Month() < time.October {
		year--
	}
	want := year*10000 + (year + 1)

	if got := formatters.GetCurrentSeasonID(); got != want {
		t.Errorf("GetCurrentSeasonID() = %v, want %v", got, want)
	}
}
