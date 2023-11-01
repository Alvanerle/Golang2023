package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Printers PrinterModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Printers: PrinterModel{DB: db},
	}
}
