package data

import (
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/utilities"
	"encoding/json"
	"strconv"
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

func (j *Job) UnmarshalJSON(data []byte) error {

	var raw struct {
		ID          json.RawMessage `json:"id"`
		Platform    string          `json:"platform"`
		Title       string          `json:"title"`
		URL         string          `json:"url"`
		ReferenceID string          `json:"referenceId"`
		PosterID    string          `json:"posterId"`
		Company     Company         `json:"company"`
		Location    string          `json:"location"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	j.ID, _ = strconv.Atoi(string(raw.ID))
	j.Platform, _ = strconv.Atoi(string(raw.Platform))
	j.Title = raw.Title
	j.URL = raw.URL
	j.ReferenceID = raw.ReferenceID
	j.PosterID = raw.PosterID
	j.Company = raw.Company
	j.Location = raw.Location

	return nil
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
