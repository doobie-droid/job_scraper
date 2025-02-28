package constants

import (
	"os"
	"strings"
	"time"
)

const MAX_OPEN_CONNECTIONS = 10
const MAX_IDLE_CONNECTIONS = 10
const DB_TIMEOUT = 3 * time.Second

// TODO change GetENV to config usage

var LINKEDIN_EMAIL = getEnv("LINKEDIN_EMAIL", "")
var LINKEDIN_PASSWORD = getEnv("LINKEDIN_PASSWORD", "")
var VALID_KEYWORDS = getEnv("VALID_KEYWORDS", "")
var RAPID_API_KEY = getEnv("RAPID_API_KEY", "")
var RAPID_API_URL = getEnv("RAPID_API_URL", "")
var DATE_POSTED string = getEnv("DATE_POSTED", "past24Hours")
var JOB_KEYWORD string = getEnv("JOB_KEYWORD", "golang")
var LOCATION string = getEnv("LOCATION", "NGA")
var LOCATION_TYPE string = getEnv("LOCATION_TYPE", "Remote")
var CITY string = getEnv("CITY", "LAGOS")

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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
