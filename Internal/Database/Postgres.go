package Database

import (
	"SiteMonitor/Internal/Config"
	"context"
	"fmt"
	"log/slog"

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

func InsertTarget(ctx context.Context, pool *pgxpool.Pool, target CreateTargetParams) error {

	queries := New(pool)
	if queries == nil {
		slog.Error("No target queries available.", "error", target)
	}

	_, err := queries.CreateTarget(ctx, target)
	if err != nil {
		slog.Error("Create target failed", "error", err)
		return err
	}

	slog.Info("Created target.", target)
	return nil
}

func GetActiveTargets(ctx context.Context, pool *pgxpool.Pool) []Target {

	queries := New(pool)
	if queries == nil {
		slog.Error("No active target queries available.", "error", nil)
	}

	activeTargets, _ := queries.ListActiveTargets(ctx)

	return activeTargets
}
