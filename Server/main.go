package main

import (
	"SiteMonitor/Internal/Database"
	_ "context"
	"fmt"
	"log/slog"
	_ "net/http"

	"SiteMonitor/Internal/Config"
	_ "SiteMonitor/Internal/Database"

	"github.com/joho/godotenv"
)

func main() {

	var env = godotenv.Load()
	if env != nil {
		slog.Warn("Error loading .env file")
	}

	var config Config.Config = Config.Load(env)
	slog.Info("Current environment: " + config.AppEnv)

	dbConnection := Database.ConnectDatabase(config)

	fmt.Println(&dbConnection)
}
