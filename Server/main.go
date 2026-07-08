package main

import (
	_ "context"
	"log/slog"
	_ "net/http"

	"SiteMonitor/Internal/Config"

	"github.com/joho/godotenv"
)

func main() {

	var env = godotenv.Load()
	if env != nil {
		slog.Warn("Error loading .env file")
	}

	var config Config.Config = Config.Load(env)

	slog.Info(config.AppEnv)

}
