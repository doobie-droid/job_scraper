package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"doobie-droid/job-scraper/utilities"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var golangProjectsJobUrl = "https://www.golangprojects.com/"

// this particular job crawler would only run if your job keyword contains go in it because the platform only lists go jobs.
// It also defaults to checking for jobs in the last month because there are no ways to filter due to time on the platform
func (platform *Platform) GolangProjects() []*data.Job {
	if !strings.Contains(strings.ToLower(platform.Cfg.JobKeyword), "go") {
		return nil
	}

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

	if strings.Contains(strings.ToLower(platform.Cfg.LocationType), "remote") {
		golangProjectsJobUrl = fmt.Sprintf("%s/golang-remote-jobs.html", golangProjectsJobUrl)
	}

	return platform.getListOfValidGolangProjectJobs(ctx)

}

func (platform *Platform) getListOfValidGolangProjectJobs(ctx context.Context) []*data.Job {
	totalJobCount := 12
	var listOfValidJobs []*data.Job
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	for index := 0; index < totalJobCount; index++ {
		jobTitleDiv := "hr.clear-both + a"
		jobUrlLink := jobTitleDiv
		jobPostingDateDiv := fmt.Sprintf("%s + i.text-sm", jobTitleDiv)
		var jobTitle, jobUrl, jobPostingDate string
		err := chromedp.Run(ctx,
			chromedp.Navigate(golangProjectsJobUrl),
			chromedp.Sleep(1*time.Second),
			chromedp.WaitVisible(jobTitleDiv),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].href`, jobUrlLink, index), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].textContent`, jobPostingDateDiv, index), &jobPostingDate),
		)

		if err != nil {
			fmt.Println("could not collect info regarding job: ", err)
		}

		IsLessThanTwoMonths, err := utilities.IsLessThanTwoMonths(jobPostingDate)
		fmt.Println(jobTitle, jobUrl, jobPostingDate, IsLessThanTwoMonths)
		if err != nil {
			fmt.Println("error checking time of job Posting:", err)
			break
		}
		if !IsLessThanTwoMonths {
			break
		}

		job := data.Job{
			Platform: data.GolangProjects,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: getCompanyName(jobTitle)},
			Location: platform.Cfg.LocationType,
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

func getCompanyName(jobTitlePlusCompanyName string) string {
	var companyName string
	arrayWithCompanyNameAsLastElement := strings.Split(jobTitlePlusCompanyName, "-")
	companyName = arrayWithCompanyNameAsLastElement[len(arrayWithCompanyNameAsLastElement)-1]
	return companyName
}
