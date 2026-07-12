package main

import (
	"SiteMonitor/Internal/Config"
	"SiteMonitor/Internal/Database"
	"context"
	"fmt"
	"log/slog"
	_ "net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var config Config.Config = Config.Load()
	slog.Info("Current environment: " + config.AppEnv)

	dbConnection := Database.ConnectDatabase(config)
	if dbConnection == nil {
		return
	}
	defer func(conn *pgxpool.Pool, ctx context.Context) {
		conn.Close()
		slog.Info("Database connection closed.")
	}(dbConnection, context.Background())

	//var target Database.CreateTargetParams = Database.CreateTargetParams{
	//	Url:             "https://www.youtube.com/",
	//	Name:            "YouTube",
	//	Environment:     config.AppEnv,
	//	IntervalSeconds: 30,
	//	TimeoutSeconds:  5,
	//}
	//err := Database.InsertTarget(context.Background(), dbConnection, target)
	//if err != nil {
	//	return
	//}

	var activeTargets = Database.GetActiveTargets(context.Background(), dbConnection)
	for index, activeTarget := range activeTargets {
		fmt.Println(index, ". ", activeTarget)
	}
}

//var targetStats Database.GetTargetStatsParams = Database.GetTargetStatsParams{
//	TargetID: pgtype.UUID(parseUUID("f1df4ff8 - f189 - 4173 - ae22 - d353e6d70885")),
//	Limit:    5,
//}

//Database.ListTargetStats(context.Background(), dbConnection, targetStats)
