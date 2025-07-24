package repository

import (
	"database/sql"

	"github.com/jonathanhu237/binding-manager/backend/internal/config"
)

type Repository struct {
	cfg    *config.Config
	dbpool *sql.DB
}

func New(cfg *config.Config, dbpool *sql.DB) *Repository {
	return &Repository{
		cfg:    cfg,
		dbpool: dbpool,
	}
}
