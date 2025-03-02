package platforms

import (
	"doobie-droid/job-scraper/config"
	"doobie-droid/job-scraper/data"
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
	// validJobs = append(validJobs, LinkedInUsingRapidApi()...)
	// validJobs = append(validJobs, LinkedInWithScraper()...)
	// validJobs = append(validJobs, TestGorilla()...)
	// validJobs = append(validJobs, Workable()...)
	// validJobs = append(validJobs, RemoteAfrica()...)
	validJobs = append(validJobs, platform.GolangProjects()...)
	return validJobs

}
