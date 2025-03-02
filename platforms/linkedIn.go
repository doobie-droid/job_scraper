package platforms

import (
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type LinkedIn struct {
}

func (platform *Platform) LinkedInUsingRapidApi() []*data.Job {
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	url := fmt.Sprintf("https://linkedin-data-api.p.rapidapi.com/search-job?%s&%s&%s",
		fmt.Sprint("keywords=", platform.Cfg.JobKeyword),
		fmt.Sprint("locationID=", platform.getLinkedInLocationId()),
		fmt.Sprint("datePosted=", platform.Cfg.DatePosted),
	)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", platform.Cfg.RapidAPIKey)
	req.Header.Add("x-rapidapi-host", platform.Cfg.RapidAPIURL)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("could not make request", err.Error())
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	resBody := string(body)

	var response data.Response[data.Job]
	if err := json.Unmarshal([]byte(resBody), &response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	var newJobs []*data.Job
	for _, job := range response.Data {
		job.Platform = data.LinkedIn
		if !jobRepo.Exists(&job) {
			newJobs = append(newJobs, &job)
			jobRepo.InsertJob(&job)
		}
	}

	var validJobs []*data.Job
	for _, newJob := range newJobs {
		if newJob.IsValid() {
			validJobs = append(validJobs, newJob)
		}
	}

	return validJobs

}

func (platform *Platform) getLinkedInLocationId() string {
	CountryCodeToLocationIds := map[string]string{"NGA": "105365761"}
	return CountryCodeToLocationIds[strings.ToUpper(platform.Cfg.Location)]

}

func (platform *Platform) getLinkedInLocationType() string {
	LocationTypeToIds := map[string]string{"remote": "2", "on-site": "1", "hybrid": "3"}
	return LocationTypeToIds[strings.ToLower(platform.Cfg.LocationType)]

}

func (platform *Platform) getLinkedInDurationCode() string {
	DurationOfPostingToDurationCode := map[string]string{"past24hours": "r86400", "pastweek": "r604800", "pastmonth": "r2592000"}
	return DurationOfPostingToDurationCode[strings.ToLower(platform.Cfg.DatePosted)]

}
