package utilities

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
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
