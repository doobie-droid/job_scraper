package data

import (
	"testing"
)

func TestJobIsValid(t *testing.T) {
	testCases := []struct {
		jobTitle        string
		ExpectedOutcome bool
	}{
		{"Go engineer", true},
		{"Backend Doctor", true},
		{"Pest Terminator", false},
		{"Pastor", false},
	}

	for _, testCase := range testCases {
		job := new(Job)
		job.Title = testCase.jobTitle
		jobIsValid := job.IsValid()
		if jobIsValid != testCase.ExpectedOutcome {
			t.Errorf("Expected job is valid to be %t, Got %t", testCase.ExpectedOutcome, jobIsValid)
		}
	}

}
