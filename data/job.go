package data

import (
	"doobie-droid/job-scraper/config"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	LinkedIn = iota
	WeWorkRemotely
	TestGorilla
	Workable
	RemoteAfrica
	GolangProjects
	BreezyHr
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
	if job.Platform == TestGorilla {
		urlAsArray := strings.Split(job.URL, "jobs/")
		stringContainingId := urlAsArray[1]
		stringContainingIdAsArray := strings.Split(stringContainingId, "-")
		return stringContainingIdAsArray[0]
	}
	if job.Platform == Workable {
		urlAsArray := strings.Split(job.URL, "view/")
		stringContainingId := urlAsArray[1]
		stringContainingIdAsArray := strings.Split(stringContainingId, "/")
		return stringContainingIdAsArray[0]
	}

	if job.Platform == RemoteAfrica {
		urlAsArray := strings.Split(job.URL, "f/")
		slug := urlAsArray[1]
		return fmt.Sprintf("remote-africa-%s", slug)
	}
	if job.Platform == GolangProjects {
		urlAsArray := strings.Split(job.URL, ".com/")
		slug := urlAsArray[1]
		return fmt.Sprintf("golang-projects-%s", slug[len(slug)-20:len(slug)-1])
	}
	return ""
}

func (job *Job) IsValid() bool {
	validKeywords := config.NewConfig().ValidKeywords
	validKeywordsArray := strings.Split(validKeywords, ",")
	for _, validJobKeyword := range validKeywordsArray {
		if strings.Contains(strings.ToLower(job.Title), strings.ToLower(validJobKeyword)) {
			return true
		}
	}
	return false
}

func (job *Job) IsValidLocation(locationString string) bool {
	conf := config.NewConfig()
	validLocations := conf.ValidLocations
	validLocationsArray := strings.Split(validLocations, ",")
	validLocationsArray = append(validLocationsArray, "worldwide", conf.City, conf.Location)
	for _, validLocation := range validLocationsArray {
		if strings.Contains(strings.ToLower(locationString), strings.ToLower(validLocation)) {
			return true
		}
	}
	return false
}
