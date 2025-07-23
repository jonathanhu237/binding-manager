package domain

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
	Version      int    `json:"version"`
}
