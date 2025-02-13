package job

import (
	"database/sql"
	"doobie-droid/job-scraper/repository"
	"log"
)

type Job struct {
	DB *sql.DB
}

func NewJobConnection() *Job {
	connection, err := repository.CreateDBConnection()
	if err != nil {
		log.Fatal("Could not connect to the db")
		log.Fatal(err)
	}
	return &Job{
		DB: connection,
	}
}
