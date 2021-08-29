package env

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DevEnv  = "development"
	TestEnv = "test"
	ProdEnv = "production"
)

func Load() {

	env := os.Getenv("KNITTO_APP_ENV")

	if env == "" {
		env = DevEnv
	}

	err := godotenv.Load(".env." + env + ".local")

	if err != nil {
		err = godotenv.Load(".env." + env)
	}

	if err != nil {
		err = godotenv.Load(".env.local")
	}

	if err != nil {
		godotenv.Load()
	}
}

func Get(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
