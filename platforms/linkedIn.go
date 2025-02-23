package platforms

import (
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"doobie-droid/job-scraper/utilities"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func LinkedInUsingRapidApi() []*data.Job {
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	url := fmt.Sprintf("https://linkedin-data-api.p.rapidapi.com/search-jobs?keywords=%s&locationId=%s&datePosted=%s", constants.JOB_KEYWORD, constants.RAPID_API_LINKEDIN_LOCATION_ID, constants.RAPID_DATE_POSTED)
	// url := "http://localhost:8000/api/v1/ping"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", utilities.GetEnv(constants.RAPID_API_ENV_KEY))
	req.Header.Add("x-rapidapi-host", utilities.GetEnv(constants.RAPID_API_ENV_URL))

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
