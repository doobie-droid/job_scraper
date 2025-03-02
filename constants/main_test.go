package constants

import "testing"

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
		LOCATION_TYPE = tc.input
		got := GetLinkedInLocationType()
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
		LOCATION = tc.input
		got := GetLinkedInLocationId()
		if got != tc.expected {
			t.Errorf("GetLinkedInLocationId(%s) expected = %s; but got %s", tc.input, got, tc.expected)
		}
	}
}

func TestGetWorkableDurationCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"past24Hours", "1"},
		{"pastWeek", "7"},
		{"pastMonth", "30"},
		{"invalid", ""},
	}

	for _, tc := range tests {
		DATE_POSTED = tc.input
		got := GetWorkableDurationCode()
		if got != tc.expected {
			t.Errorf("GetWorkableDurationCode(%s) expected = %s; but got %s", tc.input, got, tc.expected)

		}
	}
}
