package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var pictureAvatarDisplay = "#ember17"

func (platform *Platform) LinkedInWithScraper() []*data.Job {
	var linkedInJobUrl = fmt.Sprintf("https://www.linkedin.com/jobs/search/?%s&%s&%s&%s",
		fmt.Sprint("f_WT=", platform.getLinkedInLocationType()),
		fmt.Sprint("f_TPR=", platform.getLinkedInDurationCode()),
		fmt.Sprint("geoId=", platform.getLinkedInLocationId()),
		fmt.Sprint("keywords=", platform.Cfg.JobKeyword),
	)
	userDataDir := "./chromedp-profile"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
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

	return getListOfValidJobs(countOfAvailableJobs, ctx)
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
	// TODO: Need to add way of checking if logged in details could not load, since this would stall indefinitely
	var availableJobsElement = "div.jobs-search-results-list__subtitle span"
	var availableJobs string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(10*time.Second),
		chromedp.WaitVisible(pictureAvatarDisplay),
		chromedp.Text(availableJobsElement, &availableJobs),
	)
	if err != nil {
		return 0, err
	}

	return getCount(availableJobs)
}

func getListOfValidJobs(countOfAvailableJobs int, ctx context.Context) []*data.Job {
	var closeJobButton = "button.job-card-container__action.job-card-container__action-small.artdeco-button.artdeco-button--muted.artdeco-button--2.artdeco-button--tertiary.ember-view"
	var jobDetailsDiv = "div.jobs-search__job-details--wrapper"
	var jobTitleDiv = "h1.t-24.t-bold.inline"
	var jobCompanyDiv = "div.job-details-jobs-unified-top-card__company-name a.kvxOIgzTihjbTHaLGrgQguzNnXGpbPhmw"
	var jobLocationDiv = "span.tvm__text.tvm__text--low-emphasis"
	var jobDetails, jobTitle, jobCompany, jobLocation string
	var firstJobDiv = "li.ember-view"
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	var currentURL string
	var listOfValidJobs []*data.Job
	for range countOfAvailableJobs {
		err := chromedp.Run(ctx,
			chromedp.Sleep(10*time.Second),
			chromedp.Click(firstJobDiv),
			chromedp.Sleep(2*time.Second),
			chromedp.Text(jobDetailsDiv, &jobDetails),
			chromedp.WaitVisible(closeJobButton),
			chromedp.Click(closeJobButton, chromedp.NodeVisible),
			chromedp.Text(jobTitleDiv, &jobTitle),
			chromedp.Text(jobCompanyDiv, &jobCompany),
			chromedp.Text(jobLocationDiv, &jobLocation),
			chromedp.Sleep(1*time.Second),
			chromedp.Location(&currentURL),
			chromedp.Reload(),
			chromedp.WaitVisible(pictureAvatarDisplay),
		)
		if err != nil {
			fmt.Println("error extracting job details", err)
		}
		job := data.Job{
			Platform: data.LinkedIn,
			Title:    jobTitle,
			Company:  data.Company{Name: jobCompany},
			URL:      getValidUrl(currentURL),
			Location: jobLocation,
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

// converts a url that looks like https://www.linkedin.com/jobs/search/?currentJobId=4157968171&f_TPR=r86400&f_WT=2&geoId=105365761
// to a clean url that looks like https://www.linkedin.com/jobs/view/4157968171
func getValidUrl(invalidUrl string) (validUrl string) {
	urlAsArray := strings.Split(invalidUrl, "=")
	urlPartContainingId := urlAsArray[1]
	urlPartContainingIdAsArray := strings.Split(urlPartContainingId, "&")
	Id := urlPartContainingIdAsArray[0]
	return fmt.Sprintf("https://www.linkedin.com/jobs/view/%s", Id)
}
