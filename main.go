package main

import (
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/platforms"
)

func main() {
	validJobs := []*data.Job{}
	validJobs = append(validJobs, platforms.LinkedInUsingRapidApi()...)
	_ = validJobs
}
