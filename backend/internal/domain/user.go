package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password Password  `json:"-"`
	IsAdmin  bool      `json:"is_admin"`
	Version  int       `json:"version"`
}
