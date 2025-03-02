package data

import (
	"doobie-droid/job-scraper/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	setup()

	exitCode := m.Run()

	tearDown()

	os.Exit(exitCode)
}

func setup() {
	config.EnvPath = "../.env.example"

}

func tearDown() {
	config.EnvPath = "./.env"
}
