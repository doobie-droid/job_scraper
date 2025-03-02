package repository

import (
	"database/sql"
	"doobie-droid/job-scraper/config"
	"doobie-droid/job-scraper/constants"
	"fmt"

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

	db.SetConnMaxLifetime(constants.DB_TIMEOUT)
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
	userName := config.GetEnv("DB_USERNAME")
	password := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")

	return fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
}
