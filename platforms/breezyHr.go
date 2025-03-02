package platforms

import (
	"context"
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
)

func (platform *Platform) BreezyHr() []*data.Job {
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	url := fmt.Sprintf("%s?%s&%s&%s",
		constants.GOOGLE_SEARCH_API_URL,
		fmt.Sprint("key=", platform.Cfg.GoogleApiKey),
		fmt.Sprint("cx=", platform.Cfg.SearchEngineId),
		"q=site:breezy.hr",
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
	var listOfValidJobs []*data.Job
	for _, result := range response.Items {
		if result.IsValidJobListing() {
			listOfValidJobs = append(listOfValidJobs, platform.extractValidJobs(result.Link)...)
		}
	}

	return listOfValidJobs
}

func (platform *Platform) extractValidJobs(jobUrl string) []*data.Job {
	userDataDir := "./chromedp-profile"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
		chromedp.Flag("headless", false),
		chromedp.Flag("start-maximized", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	for index := range getCountOfJobElements(jobUrl, ctx) {
		jobTitleDiv := "li.position-details h2"
		jobUrlLink := "li.position-details a"
		locationDiv := "li.location"
		companyNameDiv := "a.brand img"
		var jobTitle, jobUrl, companyName, location string
		err := chromedp.Run(ctx,
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].href", jobUrlLink, index), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf("document.querySelector('%s').alt", companyNameDiv), &companyName),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", locationDiv, index), &location),
		)
		job := data.Job{
			Platform: data.BreezyHr,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: companyName},
			Location: location,
		}
		fmt.Println(job)
		if err != nil {
			fmt.Println("could not retrieve job details", err)
		}
	}

	return nil
}

func getCountOfJobElements(jobUrl string, ctx context.Context) int {
	jobTitleDiv := "li.position-details"
	var jobCount int
	err := chromedp.Run(ctx,
		chromedp.Navigate(jobUrl),
		chromedp.WaitVisible("img", chromedp.ByQuery),
		chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s').length", jobTitleDiv), &jobCount),
	)
	if err != nil {
		fmt.Println("could not get Count of Job elements", err, jobUrl)
	}
	return jobCount
}
