package apiserver

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jonathanhu237/binding-manager/backend/internal/config"
	"github.com/jonathanhu237/binding-manager/backend/internal/repository"
)

type ApiServer struct {
	logger *slog.Logger
	cfg    *config.Config
	repo   *repository.Repository
}

func New(logger *slog.Logger, cfg *config.Config, repo *repository.Repository) *ApiServer {
	return &ApiServer{
		logger: logger,
		cfg:    cfg,
		repo:   repo,
	}
}

func (as *ApiServer) Start() {
	if err := as.init(); err != nil {
		as.logger.Error(err.Error())
		return
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", as.cfg.ApiServer.Port),
		Handler: as.routes(),
	}

	as.logger.Info("Starting API server", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		as.logger.Error(err.Error())
	}
}
