package platforms

import (
	"testing"
)

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
		platform.Cfg.DatePosted = tc.input
		got := platform.getWorkableDurationCode()
		if got != tc.expected {
			t.Errorf("GetWorkableDurationCode(%s) expected = %s; but got %s", tc.input, got, tc.expected)

		}
	}
}
