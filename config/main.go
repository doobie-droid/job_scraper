package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvPath = "./.env"

type Config struct {
	LinkedinEmail    string
	LinkedinPassword string
	ValidKeywords    string
	RapidAPIKey      string
	DatePosted       string
	JobKeyword       string
	Location         string
	LocationType     string
	City             string
	ValidLocations   string
	GoogleApiKey     string
	SearchEngineId   string
}

func NewConfig() *Config {

	return &Config{
		LinkedinEmail:    GetEnv("LINKEDIN_EMAIL"),
		LinkedinPassword: GetEnv("LINKEDIN_PASSWORD"),
		ValidKeywords:    GetEnv("VALID_KEYWORDS"),
		RapidAPIKey:      GetEnv("RAPID_API_KEY"),
		DatePosted:       GetEnvOrUseDefault("DATE_POSTED", "past24Hours"),
		JobKeyword:       GetEnvOrUseDefault("JOB_KEYWORD", "golang"),
		Location:         GetEnvOrUseDefault("LOCATION", "NGA"),
		LocationType:     GetEnvOrUseDefault("LOCATION_TYPE", "Remote"),
		City:             GetEnvOrUseDefault("CITY", "LAGOS"),
		ValidLocations:   GetEnvOrUseDefault("LOCATION_KEYWORDS", "Africa, Nigeria, Worldwide"),
		GoogleApiKey:     GetEnv("GOOGLE_API_KEY"),
		SearchEngineId:   GetEnv("SEARCH_ENGINE_ID"),
	}
}

func GetEnv(key string) string {
	err := godotenv.Load(EnvPath)
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	if os.Getenv(key) == "" {
		panic(fmt.Sprintf("pls load .env key %s in the .env file", key))
	}
	return os.Getenv(key)
}

func GetEnvOrUseDefault(key string, defaultValue string) string {
	envValue := GetEnv(key)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}
