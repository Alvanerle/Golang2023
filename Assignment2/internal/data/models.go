package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Printers PrinterModel
	Tokens   TokenModel // Add a new Tokens field.
	Users    UserModel  // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Printers: PrinterModel{DB: db},
		Tokens:   TokenModel{DB: db}, // Initialize a new TokenModel instance.
		Users:    UserModel{DB: db},  // Initialize a new UserModel instance.
	}
}
