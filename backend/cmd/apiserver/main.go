package main

import (
	"log/slog"
	"os"

	"github.com/jonathanhu237/binding-manager/backend/internal/apiserver"
	"github.com/jonathanhu237/binding-manager/backend/internal/config"
)

func main() {
	// Initialize the logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load the configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration.", "error", err)
		os.Exit(1)
	}

	// Create a new API server instance.
	apiServer := apiserver.New(logger, cfg)

	// Start the API server.
	apiServer.Start()
}
