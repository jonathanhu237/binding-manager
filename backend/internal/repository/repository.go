package repository

import (
	"database/sql"

	"github.com/jonathanhu237/binding-manager/backend/internal/config"
	"github.com/jonathanhu237/binding-manager/backend/internal/domain"
)

type Repository struct {
	User interface {
		CheckAdminExists() (bool, error)
		Insert(*domain.User) error
		GetByUsername(username string) (*domain.User, error)
	}
}

func New(cfg *config.Config, dbpool *sql.DB) *Repository {
	return &Repository{
		User: &UserRepository{cfg, dbpool},
	}
}
