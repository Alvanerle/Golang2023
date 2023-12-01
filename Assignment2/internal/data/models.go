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
	Printers    PrinterModel
	Permissions PermissionModel // Add a new Permissions field.
	Tokens      TokenModel      // Add a new Tokens field.
	Users       UserModel       // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Printers:    PrinterModel{DB: db},
		Permissions: PermissionModel{DB: db}, // Initialize a new PermissionModel instance.
		Tokens:      TokenModel{DB: db},      // Initialize a new TokenModel instance.
		Users:       UserModel{DB: db},       // Initialize a new UserModel instance.
	}
}
