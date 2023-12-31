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
	Movies      MovieModel
	Permissions PermissionModel
	Users       UserModel // Add a new Users field.
	Tokens      TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Users:       UserModel{DB: db}, // Initialize a new UserModel instance.
		Tokens:      TokenModel{DB: db},
	}
}
