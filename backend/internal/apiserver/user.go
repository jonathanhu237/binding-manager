package apiserver

import (
	"net/http"

	"github.com/jonathanhu237/binding-manager/backend/internal/domain"
	"github.com/jonathanhu237/binding-manager/backend/internal/unierror"
	"github.com/jonathanhu237/binding-manager/backend/internal/validator"
)

func (as *ApiServer) createUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}

	if err := as.readJSON(w, r, &payload); err != nil {
		as.badRequestError(w, r, err)
		return
	}

	v := validator.New()

	v.Check(len(payload.Username) >= 3, "username", "must be at least 3 characters long")
	v.Check(len(payload.Username) <= 15, "username", "must be at most 15 characters long")

	v.Check(len(payload.Password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(payload.Password) <= 30, "password", "must be at most 30 characters long")
	v.Check(validator.ValidatePasswordPattern(payload.Password), "password", "must contain at least one uppercase letter, one lowercase letter, one digit, and one special character")

	if !v.Valid() {
		as.badRequestError(w, r, &v.Errors)
		return
	}

	passwordHash, err := domain.GeneratePasswordHash(payload.Password)
	if err != nil {
		as.internalServerError(w, r, err)
		return
	}

	user := &domain.User{
		Username:     payload.Username,
		PasswordHash: passwordHash,
		IsAdmin:      payload.IsAdmin,
	}

	if err := as.repo.User.Insert(user); err != nil {
		switch err {
		case unierror.ErrUsernameAlreadyExists:
			as.conflictError(w, r, err)
		case unierror.ErrEmailAlreadyExists:
			as.conflictError(w, r, err)
		default:
			as.internalServerError(w, r, err)
		}
		return
	}

	as.successResponse(w, r, &envelope{"user": user})
}
