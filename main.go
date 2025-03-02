package main

import (
	"doobie-droid/job-scraper/platforms"
	"fmt"
)

func main() {
	validJobs := platforms.GetValidJobs()

	for _, job := range validJobs {
		fmt.Println(job.Title, job.URL)
	}
}
