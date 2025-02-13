package repository

import (
	"database/sql"
	"doobie-droid/job-scraper/constants"
	"doobie-droid/job-scraper/utilities"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * constants.DB_TIMEOUT_IN_MINUTES)
	db.SetMaxOpenConns(constants.MAX_OPEN_CONNECTIONS)
	db.SetMaxIdleConns(constants.MAX_IDLE_CONNECTIONS)

	return db, nil

}

func CreateDBConnection() (*sql.DB, error) {
	dsn := createDataSourceName()
	connection, err := openDB(dsn)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func createDataSourceName() string {
	userName := utilities.GetEnv("DB_USERNAME")
	password := utilities.GetEnv("DB_PASSWORD")
	dbName := utilities.GetEnv("DB_NAME")

	return fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
}
