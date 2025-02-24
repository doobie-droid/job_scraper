package constants

import (
	"doobie-droid/job-scraper/utilities"
	"strings"
	"time"
)

const MAX_OPEN_CONNECTIONS = 10
const MAX_IDLE_CONNECTIONS = 10
const DB_TIMEOUT = 3 * time.Second

var LINKEDIN_EMAIL = utilities.GetEnv("LINKEDIN_EMAIL")
var LINKEDIN_PASSWORD = utilities.GetEnv("LINKEDIN_PASSWORD")
var VALID_KEYWORDS = utilities.GetEnv("VALID_KEYWORDS")
var RAPID_API_KEY = utilities.GetEnv("RAPID_API_KEY")
var RAPID_API_URL = utilities.GetEnv("RAPID_API_URL")
var DATE_POSTED string = utilities.GetEnvOrUseDefault("DATE_POSTED", "past24Hours")
var JOB_KEYWORD string = utilities.GetEnvOrUseDefault("JOB_KEYWORD", "golang")
var LOCATION string = utilities.GetEnvOrUseDefault("LOCATION", "NGA")
var LOCATION_TYPE string = utilities.GetEnvOrUseDefault("LOCATION_TYPE", "Remote")
var CITY string = utilities.GetEnvOrUseDefault("CITY", "LAGOS")

func GetLinkedInDurationCode() string {
	DurationOfPostingToDurationCode := map[string]string{"past24hours": "r86400", "pastweek": "r604800", "pastmonth": "r2592000"}
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
