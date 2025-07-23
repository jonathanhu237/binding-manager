package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	Email        string    `json:"email"`
	Version      int       `json:"version"`
}
