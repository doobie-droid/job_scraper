package constants

import (
	"time"
)

const (
	MAX_OPEN_CONNECTIONS  = 10
	MAX_IDLE_CONNECTIONS  = 10
	DB_TIMEOUT            = 3 * time.Second
	RAPID_API_URL         = "linkedin-data-api.p.rapidapi.com"
	GOOGLE_SEARCH_API_URL = "https://www.googleapis.com/customsearch/v1"
)
