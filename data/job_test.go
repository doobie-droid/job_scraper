package data

import (
	"doobie-droid/job-scraper/config"
	"fmt"
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

func TestLocationIsValid(t *testing.T) {
	testCases := []struct {
		locationString  string
		ExpectedOutcome bool
	}{
		{fmt.Sprintf(GetBaseLocationString(), "EU"), false},
		{fmt.Sprintf(GetBaseLocationString(), "Worldwide"), true},
		{fmt.Sprintf(GetBaseLocationString(), "Nigeria"), true},
		{fmt.Sprintf(GetBaseLocationString(), config.NewConfig().City), true},
		{fmt.Sprintf(GetBaseLocationString(), config.NewConfig().Location), true},
	}

	for _, testCase := range testCases {
		job := new(Job)
		locationIsValid := job.IsValidLocation(testCase.locationString)
		if locationIsValid != testCase.ExpectedOutcome {
			t.Errorf("Expected location is valid to be %t, Got %t", testCase.ExpectedOutcome, locationIsValid)
		}
	}

}

func GetBaseLocationString() string {
	return `Ardent Blue Technologies\n\n Linkedin profile Location: 🇪🇺 %s Remote, 100% Remote \n\n\n Skills wanted, in addition to Golang:devops\ngo\ngolang\nphp\n`
}
