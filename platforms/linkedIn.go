package platforms

import (
	"doobie-droid/job-scraper/constants"
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
	log.Println("started collecting jobs via linkedin using rapid api")
	jobRepo := job.NewJobConnection()
	url := fmt.Sprintf("https://%s/search-jobs?%s&%s&%s",
		constants.RAPID_API_URL,
		fmt.Sprint("keywords=", platform.Cfg.JobKeyword),
		fmt.Sprint("locationID=", platform.getLinkedInLocationId()),
		fmt.Sprint("datePosted=", platform.Cfg.DatePosted),
	)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", platform.Cfg.RapidAPIKey)
	req.Header.Add("x-rapidapi-host", constants.RAPID_API_URL)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("could not make request", err.Error())
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response data.Response[data.Job]
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error decoding JSON:", err)
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
		if newJob.IsValid() && newJob.IsValidLocation(newJob.Location) {
			validJobs = append(validJobs, newJob)
		}
	}
	log.Println("done collecting jobs via linkedin using rapid api")
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
