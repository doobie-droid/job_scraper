package platforms

import (
	"context"
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"strings"
)

func (platform *Platform) BreezyHr() []*data.Job {
	var listOfValidJobs []*data.Job
	maxNumberOfSearchResult := 101
	for index := 1; index < maxNumberOfSearchResult; index += 10 {
		fmt.Println(index)
		url := fmt.Sprintf("%s?%s&%s&%s&%s",
			constants.GOOGLE_SEARCH_API_URL,
			fmt.Sprint("key=", platform.Cfg.GoogleApiKey),
			fmt.Sprint("cx=", platform.Cfg.SearchEngineId),
			fmt.Sprint("start=", index),
			"q=site:*.breezy.hr+-www",
		)
		req, _ := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("could not make request", err.Error())
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		resBody := string(body)

		var response data.GoogleResponse
		if err := json.Unmarshal([]byte(resBody), &response); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return nil
		}
		for _, result := range response.Items {
			if result.IsValidJobListing() {
				listOfValidJobs = append(listOfValidJobs, platform.extractValidJobsFromSite(result.Link)...)
			}
		}

	}

	return listOfValidJobs
}

func (platform *Platform) extractValidJobsFromSite(companyUrl string) []*data.Job {
	userDataDir := "./chromedp-profile"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	jobRepo := job.NewJobConnection()
	var listOfValidJobs []*data.Job
	for index := range getCountOfJobElements(companyUrl, ctx) {
		jobTitleDiv := "li.position-details h2"
		jobUrlLink := "li.position-details a"
		locationDiv := "li.location"
		var jobTitle, jobUrl, location string
		err := chromedp.Run(ctx,
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].href", jobUrlLink, index), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", locationDiv, index), &location),
		)

		job := data.Job{
			Platform: data.BreezyHr,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: getCompanyNameBreezyHr(companyUrl)},
			Location: location,
		}
		if err != nil {
			fmt.Println("could not retrieve job details", err)
		}
		if jobRepo.Exists(&job) {
			continue
		}
		jobRepo.InsertJob(&job)
		if !job.IsValid() {
			continue
		}
		if job.IsValidLocation(job.Location) {
			listOfValidJobs = append(listOfValidJobs, &job)
		}
	}

	return listOfValidJobs
}

func getCountOfJobElements(jobUrl string, ctx context.Context) int {
	fmt.Println(jobUrl)
	jobTitleDiv := "li.position-details"
	var jobCount int
	err := chromedp.Run(ctx,
		chromedp.Navigate(jobUrl),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s').length", jobTitleDiv), &jobCount),
	)
	if err != nil {
		fmt.Println("could not get Count of Job elements", err, jobUrl)
	}
	return jobCount
}

func getCompanyNameBreezyHr(jobUrl string) string {
	urlAsArray := strings.Split(jobUrl, ".")
	companyNameAsUrl := urlAsArray[0]
	companyName := strings.TrimPrefix(companyNameAsUrl, "https://")
	return strings.ToUpper(companyName[:1]) + companyName[1:]

}
