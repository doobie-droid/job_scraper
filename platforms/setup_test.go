package platforms

import (
	"doobie-droid/job-scraper/config"
	"os"
	"testing"
)

var platform Platform

func TestMain(m *testing.M) {

	setup()

	exitCode := m.Run()

	tearDown()

	os.Exit(exitCode)
}

func setup() {
	config.EnvPath = "../.env.example"
	platform.Cfg = *config.NewConfig()

}

func tearDown() {

}
