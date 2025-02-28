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
			t.Errorf("GetLinkedInLocationType(%s) = %s; want %s", tc.input, got, tc.expected)
		}
	}
}
