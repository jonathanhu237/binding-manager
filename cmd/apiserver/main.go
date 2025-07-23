package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jonathanhu237/binding-manager/internal/apiserver"
	"github.com/jonathanhu237/binding-manager/internal/config"
	"github.com/jonathanhu237/binding-manager/internal/repository"
)

func main() {
	// Initialize the logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load the configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Connect to the database.
	dbpool, err := sql.Open("pgx", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Db,
	))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Postgres.PingTimeoutSeconds)*time.Second)
	defer cancel()
	if err := dbpool.PingContext(ctx); err != nil {
		slog.Error(err.Error())
	}

	// Create a new repository instance.
	repo := repository.New(cfg, dbpool)

	// Create a new API server instance.
	apiServer := apiserver.New(logger, cfg, repo)

	// Start the API server.
	apiServer.Start()
}
