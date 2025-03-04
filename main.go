package main

import (
	"doobie-droid/job-scraper/platforms"
	"fmt"
)

func main() {
	validJobs := platforms.GetValidJobs()

	fmt.Println("valid Jobs")
	for _, job := range validJobs {
		fmt.Printf("%s-%s-%s\n", job.Title, job.Company, job.URL)
	}
}
