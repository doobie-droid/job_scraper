package platforms

import (
	"doobie-droid/job-scraper/config"
	"doobie-droid/job-scraper/data"
	"fmt"
)

var EnvPath = "./.env"

type Platform struct {
	Cfg config.Config
}

func GetValidJobs() []*data.Job {
	platform := &Platform{}
	config.EnvPath = EnvPath
	platform.Cfg = *config.NewConfig()
	var validJobs []*data.Job
	validJobs = append(validJobs, platform.LinkedInUsingRapidApi()...)
	validJobs = append(validJobs, platform.LinkedInWithScraper()...)
	validJobs = append(validJobs, platform.TestGorilla()...)
	validJobs = append(validJobs, platform.Workable()...)
	validJobs = append(validJobs, platform.RemoteAfrica()...)
	validJobs = append(validJobs, platform.GolangProjects()...)
	validJobs = append(validJobs, platform.BreezyHr()...)
	for _, job := range validJobs {
		fmt.Println(job.Title, job.Platform, job.URL)
	}

	//TODO ADD Package to send scraped information as email // message to whatsapp
	return validJobs

}
