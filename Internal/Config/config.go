package Config

import (
	"log/slog"
	_ "log/slog"
	"os"
	_ "os"
	_ "strconv"
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

func Load(env error) Config {
	slog.Info("Loading Config...")

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
