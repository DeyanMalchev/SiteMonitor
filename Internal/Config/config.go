package Config

import (
	"log/slog"
	"os"
	_ "strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port   string
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func Load() Config {
	slog.Info("Loading Config...")

	// Rename to 'err' because godotenv returns an error type, not an environment object
	err := godotenv.Load()
	if err != nil { // Check if an error ACTUALLY exists
		slog.Warn("No .env file found, reading from system environment", "error", err)
	}

	var conf Config = Config{
		AppEnv: os.Getenv("APP_ENV"),
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
	}

	return conf
}
