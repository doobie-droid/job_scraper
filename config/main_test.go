package config

import (
	"testing"
)

func TestConfigInitialization(t *testing.T) {

	EnvPath = "../.env.example"
	Cfg := NewConfig()

	if Cfg.LinkedinEmail != "clarencesolutions.inc@gmail.com" {
		t.Errorf("Expected LinkedinEmail to be 'clarencesolutions.inc@gmail.com', got '%s'", Cfg.LinkedinEmail)
	}

	if Cfg.LinkedinPassword != "pinkponyclub" {
		t.Errorf("Expected LinkedinPassword to be 'pinkponyclub', got '%s'", Cfg.LinkedinPassword)
	}

	if Cfg.ValidKeywords != "go,laravel,backend,frontend,vue,nuxt,back-end,front-end,back end,front end" {
		t.Errorf("Expected ValidKeywords to be 'go,laravel,backend,frontend,vue,nuxt,back-end,front-end,back end,front end', got '%s'", Cfg.ValidKeywords)
	}

	if Cfg.RapidAPIKey != "2169ce24cbmsh4b0c1jsn8370dfe5dfd5" {
		t.Errorf("Expected RapidAPIKey to be '2169ce24cbmsh4b0c1jsn8370dfe5dfd5', got '%s'", Cfg.RapidAPIKey)
	}

	if Cfg.DatePosted != "past24Hours" {
		t.Errorf("Expected DatePosted to be 'past24Hours', got '%s'", Cfg.DatePosted)
	}

	if Cfg.JobKeyword != "golang" {
		t.Errorf("Expected JobKeyword to be 'golang', got '%s'", Cfg.JobKeyword)
	}

	if Cfg.Location != "NGA" {
		t.Errorf("Expected Location to be 'NGA', got '%s'", Cfg.Location)
	}

	if Cfg.LocationType != "Remote" {
		t.Errorf("Expected LocationType to be 'Remote', got '%s'", Cfg.LocationType)
	}

	if Cfg.City != "LAGOS" {
		t.Errorf("Expected City to be 'LAGOS', got '%s'", Cfg.City)
	}

	if Cfg.ValidLocations != "Africa, Nigeria" {
		t.Errorf("Expected Locations to be 'Africa, Nigeria', got '%s'", Cfg.ValidLocations)
	}

	if Cfg.GoogleApiKey != "GOOGLE_API_KEY_VALUE" {
		t.Errorf("Expected Api Key to be 'GOOGLE_API_KEY_VALUE', got %s", Cfg.GoogleApiKey)
	}

	if Cfg.SearchEngineId != "SEARCH_ENGINE_ID_VALUE" {
		t.Errorf("Expected SEARCH ENGINE ID to be 'SEARCH_ENGINE_ID_VALUE', got %s", Cfg.SearchEngineId)
	}
}
