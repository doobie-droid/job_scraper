package platforms

import (
	"doobie-droid/job-scraper/data"
	"fmt"
)

var workableJobUrl = fmt.Sprintf("https://jobs.workable.com/search?location=Lagos&day_range=30&workplace=hybrid&workplace=remote")

func Workable() []*data.Job {
	return nil
}
