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
	Users    UserModel // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Printers: PrinterModel{DB: db},
		Users:    UserModel{DB: db}, // Initialize a new UserModel instance.
	}
}
