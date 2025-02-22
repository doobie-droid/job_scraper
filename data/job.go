package data

import (
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/utilities"
	"strings"
)

const (
	LinkedIn = iota
	WeWorkRemotely
)

type Job struct {
	ID          int     `json:"id"`
	Platform    int     `json:"platform"`
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	ReferenceID string  `json:"referenceId"`
	PosterID    string  `json:"posterId"`
	Company     Company `json:"company"`
	Location    string  `json:"location"`
}

func (job *Job) GetSlug() string {
	if job.Platform == LinkedIn {
		urlAsArray := strings.Split(job.URL, "/")
		jobId := urlAsArray[len(urlAsArray)-1]
		return jobId
	}
	return ""
}

func (job *Job) IsValid() bool {
	validKeywords := utilities.GetEnv(constants.VALID_KEYWORDS_ENV_KEY)
	validKeywordsArray := strings.Split(validKeywords, ",")
	for _, validJobKeyword := range validKeywordsArray {
		if strings.Contains(strings.ToLower(job.Title), strings.ToLower(validJobKeyword)) {
			return true
		}
	}
	return false
}
