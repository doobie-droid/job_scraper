package constants

import (
	// "doobie-droid/job-scraper/config"
	// "strings"
	"time"
)

const (
	MAX_OPEN_CONNECTIONS = 10
	MAX_IDLE_CONNECTIONS = 10
	DB_TIMEOUT           = 3 * time.Second
	GOOGLE_API_URL       = "https://www.googleapis.com/customsearch/v1?"
	RAPID_API_URL        = "linkedin-data-api.p.rapidapi.com"
)
