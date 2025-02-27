package main

import (
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/platforms"
	"fmt"
)

func main() {
	validJobs := []*data.Job{}
	// validJobs = append(validJobs, platforms.LinkedInUsingRapidApi()...)
	// validJobs = append(validJobs, platforms.LinkedInWithScraper()...)
	// validJobs = append(validJobs, platforms.TestGorilla()...)
	// validJobs = append(validJobs, platforms.Workable()...)
	validJobs = append(validJobs, platforms.RemoteAfrica()...)
	fmt.Println("valid jobs")
	for _, job := range validJobs {
		fmt.Println(job.Title, job.URL)
	}
}
