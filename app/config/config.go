// file to contain server config and loading env file.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Env            string
	DBUrl          string
	RollbarToken   string
	RollbarCodeVer string
	RollbarHost    string
	RollbarRoot    string
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required but not set", key))
	}
	return value
}
func loadConfig() config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	var config config
	config.Env = GetEnv("ENV")

	dbUser := GetEnv("DB_USER")
	dbPass := GetEnv("DB_PASSWORD")
	dbHost := GetEnv("DB_HOST")
	dbPort := GetEnv("DB_PORT")
	dbName := GetEnv("DB_NAME")
	config.DBUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	config.RollbarToken = GetEnv("ROLLBAR_TOKEN")
	config.RollbarCodeVer = GetEnv("ROLLBAR_CODE_VERSION")
	config.RollbarHost = GetEnv("ROLLBAR_SERVER_HOST")
	config.RollbarRoot = GetEnv("ROLLBAR_SERVER_ROOT")

	return config
}

var Config config = loadConfig()
