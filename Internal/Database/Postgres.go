package Database

import (
	"SiteMonitor/Internal/Config"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase(config Config.Config) *pgxpool.Pool {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		slog.Error("Database connection error", err)
		pool.Close()
		return nil
	}

	if err := pool.Ping(context.Background()); err != nil {
		slog.Error("Database ping failed", "error", err)
		return nil
	}

	slog.Info("Connected to database pool.")

	return pool
}

func InsertTarget(ctx context.Context, conn *pgx.Conn, target CreateTargetParams) error {

	queries := New(conn)
	if queries == nil {
		slog.Warn("No target queries available")
	}

	_, err := queries.CreateTarget(ctx, target)
	if err != nil {
		return err
	}

	slog.Info("Created target.", target)
	return nil
}

func ListTargetStats(ctx context.Context, conn *pgx.Conn, target GetTargetStatsParams) []GetTargetStatsRow {

	queries := New(conn)
	if queries == nil {
		slog.Warn("No target queries available")
	}

	stats, err := queries.GetTargetStats(ctx, target)
	if err != nil {
		slog.Error("Error getting target stats: ", err)
	}
	slog.Info("Get target stats", target)

	return stats
}
