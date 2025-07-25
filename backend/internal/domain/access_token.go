package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

type AccessToken string

func GenerateAccessToken(secret string, user *User, expires time.Time) (AccessToken, error) {
	claims := CustomClaims{
		user.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			Subject:   user.Id.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return AccessToken(ss), nil
}
