package utilities

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(getRootEnvPath())
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

func getRootEnvPath() string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
	}

	normalizedPath := filepath.ToSlash(workingDirectory)

	pathAsArray := strings.Split(normalizedPath, "job_scraper")
	rootPath := pathAsArray[0]

	envPath := filepath.Join(rootPath, "/job_scraper/.env")
	return envPath
}
