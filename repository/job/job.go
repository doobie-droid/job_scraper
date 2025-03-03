package job

import (
	"context"
	"database/sql"
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/data"
	"doobie-droid/job-scraper/repository"
	"fmt"
	"log"
	"strings"
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

func (jobRepo *Job) Exists(job *data.Job) bool {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DB_TIMEOUT)
	defer cancel()

	query := `
		select 
			j.id
		from 
			jobs j
		where 
		    j.slug = ?`

	row := jobRepo.DB.QueryRowContext(ctx, query, job.GetSlug())

	err := row.Scan(
		&job.ID,
	)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return false
		}
		fmt.Println("error in checking job exists;", err)
	}

	return job.ID != 0
}

func (jobRepo *Job) InsertJob(job *data.Job) (*data.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DB_TIMEOUT)
	defer cancel()

	var newID int64
	stmt := `insert into jobs (url, slug, title, company, location) values (?, ?, ?, ?, ?)`

	result, err := jobRepo.DB.ExecContext(ctx, stmt,
		job.URL,
		job.GetSlug(),
		job.Title,
		job.Company.Name,
		job.Location,
	)
	if err != nil {
		//# Todo: add error handling
		fmt.Println("error inserting job:", err)
	}

	newID, err = result.LastInsertId()
	if err != nil {
		//# Todo: add error handling
		fmt.Println("error getting last inserted job:", err)
		return nil, err
	}

	job.ID = int(newID)
	return job, nil
}

// // InsertUser inserts a new user into the database, and returns the ID of the newly inserted row

// 	stmt := `insert into users (email, first_name, last_name, password, is_admin, created_at, updated_at)
// 		values ($1, $2, $3, $4, $5, $6, $7) returning id`

// 	err = m.DB.QueryRowContext(ctx, stmt,
// 		user.Email,
// 		user.FirstName,
// 		user.LastName,
// 		hashedPassword,
// 		user.IsAdmin,
// 		time.Now(),
// 		time.Now(),
// 	).Scan(&newID)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return newID, nil
// }
