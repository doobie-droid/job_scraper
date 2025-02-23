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
)

func LinkedInUsingRapidApi() []*data.Job {
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	url := fmt.Sprintf("https://linkedin-data-api.p.rapidapi.com/search-job?%s&%s&%s",
		fmt.Sprint("keywords=", constants.JOB_KEYWORD),
		fmt.Sprint("locationID=", constants.GetLocationId()),
		fmt.Sprint("datePosted=", constants.DATE_POSTED),
	)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", constants.RAPID_API_KEY)
	req.Header.Add("x-rapidapi-host", constants.RAPID_API_URL)

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
