package config

import (
	"doobie-droid/job-scraper/utilities"
)

var Cfg *Config

type Config struct {
	LinkedinEmail    string
	LinkedinPassword string
	ValidKeywords    string
	RapidAPIKey      string
	RapidAPIURL      string
	DatePosted       string
	JobKeyword       string
	Location         string
	LocationType     string
	City             string
}

func init() {
	Cfg = &Config{
		LinkedinEmail:    utilities.GetEnv("LINKEDIN_EMAIL"),
		LinkedinPassword: utilities.GetEnv("LINKEDIN_PASSWORD"),
		ValidKeywords:    utilities.GetEnv("VALID_KEYWORDS"),
		RapidAPIKey:      utilities.GetEnv("RAPID_API_KEY"),
		RapidAPIURL:      utilities.GetEnv("RAPID_API_URL"),
		DatePosted:       utilities.GetEnvOrUseDefault("DATE_POSTED", "past24Hours"),
		JobKeyword:       utilities.GetEnvOrUseDefault("JOB_KEYWORD", "golang"),
		Location:         utilities.GetEnvOrUseDefault("LOCATION", "NGA"),
		LocationType:     utilities.GetEnvOrUseDefault("LOCATION_TYPE", "Remote"),
		City:             utilities.GetEnvOrUseDefault("CITY", "LAGOS"),
	}
}
