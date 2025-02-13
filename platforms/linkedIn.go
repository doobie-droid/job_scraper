package platforms

import (
	"doobie-droid/job-scraper/repository/job"
)

// "doobie-droid/job-scraper/constants"
// "doobie-droid/job-scraper/utilities"
// "fmt"
// "io"
// "log"
// "net/http"

func LinkedIn() {
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	// fmt.Println(utilities.GetEnv(constants.RAPID_API_ENV_KEY))
	// url := fmt.Sprintf("https://linkedin-data-api.p.rapidapi.com/search-jobs?keywords=%s&locationId=%s&datePosted=%s&sort=mostRelevant", constants.JOB_KEYWORD, constants.RAPID_API_LINKEDIN_LOCATION_ID, constants.RAPID_DATE_POSTED)
	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("x-rapidapi-key", utilities.GetEnv(constants.RAPID_API_ENV_KEY))
	// req.Header.Add("x-rapidapi-host", utilities.GetEnv(constants.RAPID_API_ENV_URL))

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal("could not make request", err.Error())
	// }

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

}
