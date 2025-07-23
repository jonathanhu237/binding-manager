package apiserver

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jonathanhu237/binding-manager/backend/internal/config"
)

type ApiServer struct {
	logger *slog.Logger
	config *config.Config
}

func New(logger *slog.Logger, config *config.Config) *ApiServer {
	return &ApiServer{
		logger: logger,
		config: config,
	}
}

func (as *ApiServer) Start() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", as.config.Server.Port),
		Handler: as.routes(),
	}

	as.logger.Info("Starting API server", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		as.logger.Error(err.Error())
	}
}
