package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password Password  `json:"-"`
	Email    string    `json:"email"`
	IsAdmin  bool      `json:"is_admin"`
	Version  int       `json:"version"`
}
