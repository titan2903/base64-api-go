package utils

import (
	"os"
)

func GoDotEnvVariable(key string) string {
	// load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	return os.Getenv(key)
}
