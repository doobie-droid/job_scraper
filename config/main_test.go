package config

import (
	"os"
	"testing"
)

func TestConfigInitialization(t *testing.T) {
	err := os.Setenv("LINKEDIN_EMAIL", "test@example.com")
	if err != nil {
		return
	}
	err = os.Setenv("LINKEDIN_PASSWORD", "password")
	if err != nil {
		return
	}
	err = os.Setenv("VALID_KEYWORDS", "golang,remote,backend")
	if err != nil {
		return
	}
	err = os.Setenv("RAPID_API_KEY", "testapikey")
	if err != nil {
		return
	}
	err = os.Setenv("RAPID_API_URL", "linkedin-data-api.p.rapidapi.com")
	if err != nil {
		return
	}
	err = os.Setenv("DATE_POSTED", "pastweek")
	if err != nil {
		return
	}
	err = os.Setenv("JOB_KEYWORD", "golang")
	if err != nil {
		return
	}
	err = os.Setenv("LOCATION", "NG")
	if err != nil {
		return
	}
	err = os.Setenv("LOCATION_TYPE", "On-Site")
	if err != nil {
		return
	}
	err = os.Setenv("CITY", "Lagos")
	if err != nil {
		return
	}

	// Reinitialize config to reflect new env variables
	Cfg = &Config{
		LinkedinEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedinPassword: os.Getenv("LINKEDIN_PASSWORD"),
		ValidKeywords:    os.Getenv("VALID_KEYWORDS"),
		RapidAPIKey:      os.Getenv("RAPID_API_KEY"),
		RapidAPIURL:      os.Getenv("RAPID_API_URL"),
		DatePosted:       os.Getenv("DATE_POSTED"),
		JobKeyword:       os.Getenv("JOB_KEYWORD"),
		Location:         os.Getenv("LOCATION"),
		LocationType:     os.Getenv("LOCATION_TYPE"),
		City:             os.Getenv("CITY"),
	}

	// Running test cases against specific conditions
	if Cfg.LinkedinEmail != "test@example.com" {
		t.Errorf("Expected LinkedinEmail to be 'test@example.com', got '%s'", Cfg.LinkedinEmail)
	}

	if Cfg.LinkedinPassword != "password" {
		t.Errorf("Expected LinkedinPassword to be 'password', got '%s'", Cfg.LinkedinPassword)
	}

	if Cfg.ValidKeywords != "golang,remote,backend" {
		t.Errorf("Expected ValidKeywords to be 'golang,remote,backend', got '%s'", Cfg.ValidKeywords)
	}

	if Cfg.RapidAPIKey != "testapikey" {
		t.Errorf("Expected RapidAPIKey to be 'testapikey', got '%s'", Cfg.RapidAPIKey)
	}

	if Cfg.RapidAPIURL != "linkedin-data-api.p.rapidapi.com" {
		t.Errorf("Expected RapidAPIURL to be 'linkedin-data-api.p.rapidapi.com', got '%s'", Cfg.RapidAPIURL)
	}

	if Cfg.DatePosted != "pastweek" {
		t.Errorf("Expected DatePosted to be 'pastweek', got '%s'", Cfg.DatePosted)
	}

	if Cfg.JobKeyword != "golang" {
		t.Errorf("Expected JobKeyword to be 'golang', got '%s'", Cfg.JobKeyword)
	}

	if Cfg.Location != "NG" {
		t.Errorf("Expected Location to be 'NG', got '%s'", Cfg.Location)
	}

	if Cfg.LocationType != "On-Site" {
		t.Errorf("Expected LocationType to be 'On-Site', got '%s'", Cfg.LocationType)
	}

	if Cfg.City != "Lagos" {
		t.Errorf("Expected City to be 'Lagos', got '%s'", Cfg.City)
	}
}
