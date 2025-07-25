package apiserver

import (
	"net/http"
	"time"

	"github.com/jonathanhu237/binding-manager/backend/internal/domain"
	"github.com/jonathanhu237/binding-manager/backend/internal/unierror"
	"github.com/jonathanhu237/binding-manager/backend/internal/validator"
)

func (as *ApiServer) createAccessToken(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := as.readJSON(w, r, &payload); err != nil {
		as.badRequestError(w, r, err)
		return
	}

	v := validator.New()
	v.Check(payload.Username != "", "username", "must be provided")
	v.Check(payload.Password != "", "password", "must be provided")

	if !v.Valid() {
		as.badRequestError(w, r, &v.Errors)
		return
	}

	// Get the user by username.
	user, err := as.repo.User.GetByUsername(payload.Username)
	if err != nil {
		switch err {
		case unierror.ErrUsernameNotExists:
			// We should not reveal whether the username exists or not.
			as.unauthorizedError(w, r, unierror.ErrInvalidCredentials)
		default:
			as.internalServerError(w, r, err)
		}
		return
	}

	// Verify the password.
	if err := user.PasswordHash.Verify(payload.Password); err != nil {
		as.unauthorizedError(w, r, unierror.ErrInvalidCredentials)
		return
	}

	// Generate the access token.
	expires := time.Now().Add(time.Minute * time.Duration(as.cfg.Jwt.ExpireMinutes))
	access_token, err := domain.GenerateAccessToken(as.cfg.Jwt.Secret, user, expires)
	if err != nil {
		as.internalServerError(w, r, err)
		return
	}

	// Set the HTTP-only cookie.
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    string(access_token),
		Expires:  expires,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	if as.cfg.ApiServer.Environment == "development" {
		cookie.Secure = false
	}
	http.SetCookie(w, cookie)

	// Respond with success.
	as.successResponse(w, r, nil)
}
