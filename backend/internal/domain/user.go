package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash string

func GeneratePasswordHash(plaintext string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return PasswordHash(hash), nil
}

func (ph *PasswordHash) Verify(plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(*ph), []byte(plaintext))
}

type User struct {
	Id           uuid.UUID    `json:"id"`
	Username     string       `json:"username"`
	PasswordHash PasswordHash `json:"-"`
	IsAdmin      bool         `json:"is_admin"`
	Version      int          `json:"version"`
}
