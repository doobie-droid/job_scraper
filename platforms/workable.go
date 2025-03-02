package platforms

import (
	"context"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository/job"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func (platform *Platform) Workable() []*data.Job {
	var workableJobUrl = fmt.Sprintf("https://jobs.workable.com/search?%s&%s&%s",
		fmt.Sprintf("location=%s", platform.Cfg.City),
		fmt.Sprintf("day_range=%s", platform.getWorkableDurationCode()),
		fmt.Sprintf("workplace=%s", strings.ToLower(platform.Cfg.LocationType)),
	)

	userDataDir := "./chromedp-profile"
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	countOfValidJobs, err := platform.getCountOfAvailableWorkableJobs(ctx, workableJobUrl)
	if err != nil {
		fmt.Println("we could not get count of available jobs:", err)
	}
	return platform.getListOfValidWorkableJobs(countOfValidJobs, ctx)
}

func (platform *Platform) getListOfValidWorkableJobs(countOfAvailableJobs int, ctx context.Context) []*data.Job {
	listOfValidJobs := []*data.Job{}
	showMoreButton := "button.secondary__default--2ySVn"
	jobTitleDiv := "a.jobCardDetails__link--fXxEi"
	jobUrlLink := "a.jobCard__overlay--29JV5"
	companyLink := "a.companyName__link--2ntbf"
	jobRepo := job.NewJobConnection()
	_ = jobRepo
	var jobUrl, jobTitle, companyTitle string
	for index := range countOfAvailableJobs {
		if (index)%20 == 0 && index > 0 {
			err := chromedp.Run(ctx,
				chromedp.WaitReady(showMoreButton, chromedp.ByQuery),
				chromedp.ScrollIntoView(showMoreButton),
				chromedp.Sleep(3*time.Second),
				chromedp.Click(showMoreButton),
				chromedp.Sleep(2*time.Second),
			)
			if err != nil {
				fmt.Println("could not view more,", err)
			}
		}
		err := chromedp.Run(ctx,
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", jobTitleDiv, index), &jobTitle),
			chromedp.Evaluate(fmt.Sprintf(`document.querySelectorAll('%s')[%d].href`, jobUrlLink, index), &jobUrl),
			chromedp.Evaluate(fmt.Sprintf("document.querySelectorAll('%s')[%d].textContent", companyLink, index), &companyTitle),
			chromedp.Sleep(2*time.Second),
		)
		if err != nil {
			fmt.Println("could not read workable job:", err)
		}
		job := data.Job{
			Platform: data.Workable,
			Title:    jobTitle,
			URL:      jobUrl,
			Company:  data.Company{Name: companyTitle},
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

func (platform *Platform) getCountOfAvailableWorkableJobs(ctx context.Context, url string) (int, error) {
	var availableJobsElement = "div.jobsMasterDetailView__title--2bJIW"
	var availableJobs string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(10*time.Second),
		chromedp.WaitVisible(availableJobsElement),
		chromedp.Text(availableJobsElement, &availableJobs),
	)
	if err != nil {
		return 0, err
	}

	return getCount(availableJobs)
}

func (platform *Platform) getWorkableDurationCode() string {
	DurationOfPostingToDurationCode := map[string]string{"past24hours": "1", "pastweek": "7", "pastmonth": "30"}
	return DurationOfPostingToDurationCode[strings.ToLower(platform.Cfg.DatePosted)]
}
