// Config package contains env functions
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value from key ---
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
