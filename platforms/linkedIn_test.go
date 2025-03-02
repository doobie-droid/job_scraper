package platforms

import (
	// "doobie-droid/job-scraper/config"
	"testing"
)

func TestGetLinkedInLocationType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Remote", "2"},
		{"On-Site", "1"},
		{"Hybrid", "3"},
		{"Invalid", ""},
	}

	for _, tc := range tests {
		platform.Cfg.LocationType = tc.input
		got := platform.getLinkedInLocationType()
		if got != tc.expected {
			t.Errorf("GetLinkedInLocationType(%s) expected = %s; but got %s", tc.input, got, tc.expected)
		}
	}
}

func TestGetLinkedInLocationId(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NGA", "105365761"},
		{"GHA", ""},
	}

	for _, tc := range tests {
		platform.Cfg.Location = tc.input
		got := platform.getLinkedInLocationId()
		if got != tc.expected {
			t.Errorf("GetLinkedInLocationId(%s) expected = %s; but got %s", tc.input, got, tc.expected)
		}
	}
}

func TestGetLinkedInDurationCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"past24Hours", "r86400"},
		{"pastWeek", "r604800"},
		{"pastMonth", "r2592000"},
		{"invalid", ""},
	}

	for _, tc := range tests {
		platform.Cfg.DatePosted = tc.input
		got := platform.getLinkedInDurationCode()
		if got != tc.expected {
			t.Errorf("GetLinkedInDurationCode(%s) expected = %s; but got %s", tc.input, got, tc.expected)
		}
	}
}
