package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func (platform *Platform) LinkedInWithScraper() []*data.Job {
	log.Println("started collecting jobs via linkedin using crawler")
	var linkedInJobUrl = fmt.Sprintf("https://www.linkedin.com/jobs/search/?%s&%s&%s&%s",
		fmt.Sprint("f_WT=", platform.getLinkedInLocationType()),
		fmt.Sprint("f_TPR=", platform.getLinkedInDurationCode()),
		fmt.Sprint("geoId=", platform.getLinkedInLocationId()),
		fmt.Sprint("keywords=", platform.Cfg.JobKeyword),
	)
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

	if !sessionExists() {
		err := platform.signInToLinkedIn(ctx)
		if err != nil {
			log.Fatal("could not sign in", err)
		}
	}

	countOfAvailableJobs, err := getCountOfAvailableJobs(ctx, linkedInJobUrl)

	if err != nil {
		log.Fatal("could not get count of available jobs", err)
	}
	validJobs := getListOfValidJobs(countOfAvailableJobs, ctx)
	log.Println("done collecting jobs via linkedin using crawler")
	return validJobs
}

func (platform *Platform) signInToLinkedIn(ctx context.Context) error {
	var signInButton string = ".sign-in-form__sign-in-cta"
	var emailField string = "input#username"
	var passwordField string = "input#password"
	var submitButton string = ".btn__primary--large.from__button--floating"
	return chromedp.Run(ctx,
		chromedp.Navigate(`https://linkedin.com`),
		chromedp.WaitVisible(signInButton),
		chromedp.Click(signInButton, chromedp.NodeVisible),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible(emailField),
		chromedp.SendKeys(emailField, platform.Cfg.LinkedinEmail),
		chromedp.SendKeys(passwordField, platform.Cfg.LinkedinPassword),
		chromedp.Sleep(2*time.Second),
		chromedp.Click(submitButton, chromedp.NodeVisible),
		chromedp.Sleep(10*time.Second),
	)

}

func getCountOfAvailableJobs(ctx context.Context, url string) (int, error) {
	var availableJobsElement = "div.jobs-search-results-list__subtitle span"
	var availableJobs string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Text(availableJobsElement, &availableJobs),
	)
	if err != nil {
		return 0, err
	}

	return getCount(availableJobs)
}

func getListOfValidJobs(countOfAvailableJobs int, ctx context.Context) []*data.Job {
	var jobTitleDiv = "a.job-card-container__link span strong"
	var jobUrlLink = "a.job-card-container__link "
	var jobCompanyDiv = "div.artdeco-entity-lockup__subtitle span"
	var jobLocationDiv = "div.artdeco-entity-lockup__caption span"
	var jobTitle, jobCompany, jobLocation, jobUrl string
	jobRepo := job.NewJobConnection()
	var listOfValidJobs []*data.Job

	for index := range countOfAvailableJobs {
		jobsPerPage := 25
		lastJobInPage := jobsPerPage - 1
		pageIndex := index % jobsPerPage
		err := chromedp.Run(ctx,
			chromedp.WaitReady("body", chromedp.ByQuery),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, pageIndex), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].href", jobUrlLink, pageIndex), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobCompanyDiv, pageIndex), &jobCompany),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobLocationDiv, pageIndex), &jobLocation),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].scrollIntoView()", jobTitleDiv, pageIndex), nil),
			chromedp.Sleep(1*time.Second),
		)

		if err != nil {
			fmt.Println("error extracting job details", err)
		}

		job := data.Job{
			Platform: data.LinkedInCrawler,
			Title:    jobTitle,
			Company:  data.Company{Name: jobCompany},
			URL:      removeQueryParams(jobUrl),
			Location: jobLocation,
		}
		if jobRepo.Exists(&job) {
			if pageIndex == lastJobInPage {
				goToNextPage(ctx)
			}
			continue
		}
		jobRepo.InsertJob(&job)
		if job.IsValid() && job.IsValidLocation(job.Location) {
			listOfValidJobs = append(listOfValidJobs, &job)
		}
		if pageIndex == lastJobInPage {
			goToNextPage(ctx)
		}

	}
	return listOfValidJobs
}

// Gets the actual count of available jobs that meets the criteria from the text containing it
func getCount(jobResultsText string) (int, error) {
	jobResultsArray := strings.Split(jobResultsText, " ")
	jobCountAsText := jobResultsArray[0]
	jobCount, err := strconv.Atoi(jobCountAsText)
	if err != nil {
		return 0, err
	}
	return jobCount, nil

}

// check if the user has already logged in before
func sessionExists() bool {
	path := "./chromedp-profile"
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// converts a url that looks like https://www.linkedin.com/jobs/view/4165485800/?query=findme
// to https://www.linkedin.com/jobs/view/4157968171
func removeQueryParams(invalidUrl string) (validUrl string) {
	urlAsArray := strings.Split(invalidUrl, "?")
	urlPartContainingId := urlAsArray[0]
	return urlPartContainingId[:len(urlPartContainingId)-1]
}

func goToNextPage(ctx context.Context) {
	var nextButton = "button.jobs-search-pagination__button--next"
	randomNumber := rand.Intn(21)
	log.Printf("going to sleep for %d seconds to simulate human interaction\n", randomNumber)
	err := chromedp.Run(ctx,
		chromedp.ScrollIntoView(nextButton),
		chromedp.Sleep(time.Duration(randomNumber)*time.Second),
		chromedp.Click(nextButton),
	)
	if err != nil {
		log.Println("could not go to the next page in linkedin scraper", err)
	}
}
