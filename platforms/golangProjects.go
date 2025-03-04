package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"doobie-droid/job-scraper/utilities"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

var golangProjectsJobUrl = "https://www.golangprojects.com/"

// this particular job crawler would only run if your job keyword contains go in it because the platform only lists go jobs.
// It also defaults to checking for jobs in the last month because there are no ways to filter due to time on the platform
func (platform *Platform) GolangProjects() []*data.Job {
	log.Println("started collecting jobs via golang Projects using crawler")
	if !strings.Contains(strings.ToLower(platform.Cfg.JobKeyword), "go") {
		return nil
	}

	userDataDir := "./chromedp-profile"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	if strings.Contains(strings.ToLower(platform.Cfg.LocationType), "remote") {
		golangProjectsJobUrl = fmt.Sprintf("%s/golang-remote-jobs.html", golangProjectsJobUrl)
	}

	validJobs := platform.getListOfValidGolangProjectJobs(ctx)
	log.Println("done collecting jobs via golang Projects using crawler")
	return validJobs

}

func (platform *Platform) getListOfValidGolangProjectJobs(ctx context.Context) []*data.Job {
	numberOfJobsDisplayedPerInfiniteLoad := 12
	totalJobCount := numberOfJobsDisplayedPerInfiniteLoad
	var listOfValidJobs []*data.Job
	jobRepo := job.NewJobConnection()
	err := chromedp.Run(ctx,
		chromedp.Navigate(golangProjectsJobUrl))
	if err != nil {
		fmt.Println("could not visit the url Page")
	}
	for index := 0; index < totalJobCount; index++ {
		jobTitleDiv := "hr.clear-both + a"
		jobUrlLink := jobTitleDiv
		jobPostingDateDiv := fmt.Sprintf("%s + i.text-sm", jobTitleDiv)
		var jobTitle, jobUrl, jobPostingDate string
		err := chromedp.Run(ctx,

			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(jobTitleDiv),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].href`, jobUrlLink, index), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].textContent`, jobPostingDateDiv, index), &jobPostingDate),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].scrollIntoView()", jobTitleDiv, index), nil),
			chromedp.Sleep(2*time.Second),
		)

		if err != nil {
			fmt.Println("could not collect info regarding job: ", err)
		}

		IsLessThanTwoMonths, err := utilities.IsLessThanTwoMonths(jobPostingDate)
		if err != nil {
			fmt.Println("error checking time of job Posting:", err)
			break
		}
		if !IsLessThanTwoMonths {
			fmt.Println("greater than 2 months")
			break
		}

		job := data.Job{
			Platform: data.GolangProjects,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: getCompanyName(jobTitle)},
			Location: platform.Cfg.LocationType,
		}
		fmt.Println(index)
		if jobRepo.Exists(&job) {
			continue
		}
		jobRepo.InsertJob(&job)

		if !job.IsValid() {
			continue
		}

		locationString, err := GetGolangProjectJobLocationString(ctx, job.URL)
		if err != nil {
			fmt.Println("could not get location string", err)
		}
		if job.IsValidLocation(locationString) {
			listOfValidJobs = append(listOfValidJobs, &job)
		}

		if index%numberOfJobsDisplayedPerInfiniteLoad == 0 {
			chromedp.Run(ctx,
				chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s').length", jobTitleDiv), &totalJobCount),
			)
		}
	}
	return listOfValidJobs
}

// opens the main job url in a new tab from the list of all job urls
// this is necessary because the job locations are just rendered in the html body and not in any element
// so there is no way of using query selectors to retrieve the location string
func GetGolangProjectJobLocationString(ctx context.Context, jobUrl string) (string, error) {
	locationDiv := "div.p-4"
	var newTargetID target.ID
	var err error
	var location string

	err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		newTargetID, err = target.CreateTarget(jobUrl).Do(ctx)
		return err
	}))

	if err != nil {
		fmt.Println("could not create new target", err)
		return "", err
	}

	newTabCtx, cancelNewTab := chromedp.NewContext(ctx, chromedp.WithTargetID(newTargetID))
	defer cancelNewTab()

	err = chromedp.Run(newTabCtx,
		chromedp.WaitVisible(locationDiv, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(locationDiv, &location),
	)
	if err != nil {
		fmt.Println("could not carry out actions in new tab", err)
	}
	return location, nil
}

func getCompanyName(jobTitlePlusCompanyName string) string {
	var companyName string
	arrayWithCompanyNameAsLastElement := strings.Split(jobTitlePlusCompanyName, "-")
	companyName = arrayWithCompanyNameAsLastElement[len(arrayWithCompanyNameAsLastElement)-1]
	return companyName
}
