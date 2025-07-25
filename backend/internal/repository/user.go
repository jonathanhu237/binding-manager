package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jonathanhu237/binding-manager/backend/internal/config"
	"github.com/jonathanhu237/binding-manager/backend/internal/domain"
	"github.com/jonathanhu237/binding-manager/backend/internal/unierror"
)

type UserRepository struct {
	cfg    *config.Config
	dbpool *sql.DB
}

func (ur *UserRepository) CheckAdminExists() (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE is_admin = True);`

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ur.cfg.Postgres.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	var exists bool
	if err := ur.dbpool.QueryRowContext(ctx, query).Scan(&exists); err != nil {
		return false, unierror.ErrInternalServerError
	}

	return exists, nil
}

func (ur *UserRepository) Insert(user *domain.User) error {
	query := `
		INSERT INTO users (username, password_hash, is_admin)
		VALUES ($1, $2, $3)
		RETURNING id, version;
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ur.cfg.Postgres.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	if err := ur.dbpool.QueryRowContext(ctx, query, user.Username, user.PasswordHash, user.IsAdmin).Scan(&user.Id, &user.Version); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				switch pgErr.ConstraintName {
				case "users_username_key":
					return unierror.ErrUsernameAlreadyExists
				case "users_email_key":
					return unierror.ErrEmailAlreadyExists
				default:
					return err
				}
			}

			return err
		}
	}

	return nil
}

func (ur *UserRepository) GetByUsername(username string) (*domain.User, error) {
	query := `
		SELECT id, username, password_hash, is_admin, version
		FROM users
		WHERE username = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ur.cfg.Postgres.QueryTimeoutSeconds)*time.Second)
	defer cancel()

	var user domain.User

	if err := ur.dbpool.QueryRowContext(ctx, query, username).Scan(
		&user.Id,
		&user.Username,
		&user.PasswordHash,
		&user.IsAdmin,
		&user.Version,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, unierror.ErrUsernameNotExists
		}
		return nil, unierror.ErrInternalServerError
	}
	return &user, nil
}
