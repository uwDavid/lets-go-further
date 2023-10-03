package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflit    = errors.New("edit conflict")
)

// Models struct to wrap MovieModel + others
type Models struct {
	Movies MovieModel
	Users  UserModel // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db}, // Initialize a new UserModel instance.
	}
}
