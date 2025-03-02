package utilities

import (
	"fmt"
	"testing"
	"time"
)

func TestIsLessThanTwoMonths(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
		hasError bool
	}{
		{fmt.Sprintf("%s 2005", time.Now().AddDate(0, -1, 0).Month()), false, false},
		{fmt.Sprintf("%s 2025", time.Now().AddDate(0, -1, 0).Month()), true, false},
		{fmt.Sprintf("%s 2025", time.Now().AddDate(0, 1, 0).Month()), false, false},
		{fmt.Sprintf("%s 2025", time.Now().Month()), true, false},
		{"Jan 2025", false, true},
	}

	for _, tc := range tests {
		fmt.Println(tc.input)
		actual, err := IsLessThanTwoMonths(tc.input)
		if actual != tc.expected {
			t.Errorf("IsOneMonthAgo(%s) expected = %t; but got %t", tc.input, tc.expected, actual)
		}
		if tc.hasError == false && err != nil {
			t.Errorf("IsOneMonthAgo(%s) did not expect error ; but got %v", tc.input, err)
		}
	}
}
