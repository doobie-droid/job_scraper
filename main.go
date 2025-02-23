package main

import (
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/platforms"
	"fmt"
)

func main() {
	validJobs := []*data.Job{}
	validJobs = append(validJobs, platforms.LinkedInUsingRapidApi()...)
	for _, job := range validJobs {
		fmt.Println(job.Title, job.URL)
	}
}
