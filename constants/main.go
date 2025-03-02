package constants

import (
	"doobie-droid/job-scraper/config"
	"strings"
	"time"
)

const (
	MAX_OPEN_CONNECTIONS = 10
	MAX_IDLE_CONNECTIONS = 10
	DB_TIMEOUT           = 3 * time.Second
)

var (
	LINKEDIN_EMAIL    = config.Cfg.LinkedinEmail
	LINKEDIN_PASSWORD = config.Cfg.LinkedinPassword
	VALID_KEYWORDS    = config.Cfg.ValidKeywords
	RAPID_API_KEY     = config.Cfg.RapidAPIKey
	RAPID_API_URL     = config.Cfg.RapidAPIURL
	DATE_POSTED       = config.Cfg.DatePosted
	JOB_KEYWORD       = config.Cfg.JobKeyword
	LOCATION          = config.Cfg.Location
	LOCATION_TYPE     = config.Cfg.LocationType
	CITY              = config.Cfg.City
)

func GetLinkedInDurationCode() string {
	DurationOfPostingToDurationCode := map[string]string{"past24hours": "r86400", "pastweek": "r604800", "pastmonth": "r2592000"}
	return DurationOfPostingToDurationCode[strings.ToLower(DATE_POSTED)]

}

func GetWorkableDurationCode() string {
	DurationOfPostingToDurationCode := map[string]string{"past24hours": "1", "pastweek": "7", "pastmonth": "30"}
	return DurationOfPostingToDurationCode[strings.ToLower(DATE_POSTED)]
}

func GetLinkedInLocationId() string {
	CountryCodeToLocationIds := map[string]string{"NGA": "105365761"}
	return CountryCodeToLocationIds[strings.ToUpper(LOCATION)]

}

func GetLinkedInLocationType() string {
	LocationTypeToIds := map[string]string{"remote": "2", "on-site": "1", "hybrid": "3"}
	return LocationTypeToIds[strings.ToLower(LOCATION_TYPE)]

}
