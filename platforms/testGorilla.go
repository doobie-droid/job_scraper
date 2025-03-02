package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var testGorillaJobUrl = "https://www.testgorilla.com/careers"

// # Todo: add error handling that does not shut down entire app when something is wrong unless when that's a
// .env issue that means nothing can happen at all -> check all fmt.Println and log. functions

// # Todo: refactor all the platforms into their own independent packages with interfaces since you  can see
// some methods that are repeated e.g. joburl, listofvalidjobs .etc
func (platform *Platform) TestGorilla() []*data.Job {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	countOfAvailableJobs, err := getCountOfAvailableTestGorillaJobs(ctx)

	if err != nil {
		log.Fatal("could not get count of available jobs", err)
	}

	return getListOfValidTestGorillaJobs(countOfAvailableJobs, ctx)
}

func getCountOfAvailableTestGorillaJobs(ctx context.Context) (int, error) {
	var jobTitleDiv = "h3.sc-58c38bba-6.fYynjg"
	var availableJobs int
	var testGorillaLogo = "#testgorilla-logo-image"
	err := chromedp.Run(ctx,
		chromedp.Navigate(testGorillaJobUrl),
		chromedp.WaitVisible(testGorillaLogo, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.ScrollIntoView(jobTitleDiv),
		chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s').length", jobTitleDiv), &availableJobs),
		chromedp.Sleep(5*time.Second),
	)
	if err != nil {
		return 0, err
	}

	return availableJobs, nil
}

func getListOfValidTestGorillaJobs(countOfAvailableJobs int, ctx context.Context) []*data.Job {
	listOfValidJobs := []*data.Job{}
	jobTitleDiv := "h3.sc-58c38bba-6.fYynjg"
	jobUrlLink := "a.sc-58c38bba-5.hoErbF"
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	var jobUrl, jobTitle string
	for index := range countOfAvailableJobs {

		err := chromedp.Run(ctx,
			chromedp.Text(jobTitleDiv, &jobTitle),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].href`, jobUrlLink, index), &jobUrl),
			chromedp.Sleep(2*time.Second),
		)
		if err != nil {
			fmt.Println("could not read job:", err)
		}
		job := data.Job{
			Platform: data.TestGorilla,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: "Test Gorilla"},
			Location: "Remote",
		}
		if jobRepo.Exists(&job) {
			continue
		}
		jobRepo.InsertJob(&job)
		if job.IsValid() {
			listOfValidJobs = append(listOfValidJobs, &job)
		}

	}

	return listOfValidJobs
}
