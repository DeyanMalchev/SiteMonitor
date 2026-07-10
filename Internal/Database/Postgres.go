package Database

import (
	"SiteMonitor/Internal/Config"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDatabase(config Config.Config) *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		connClose := conn.Close(ctx)
		if connClose != nil {
			slog.Warn("Could not close database connection: %v", connClose)
			os.Exit(1)
		}
	}(conn, context.Background())

	slog.Info("Connected to database.", conn.PgConn())

	return conn
}
