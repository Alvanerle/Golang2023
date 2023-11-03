package data

import (
	"Printers.imangalizhumash.net/internal/validator"
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
)

type Printer struct {
	ID                  int64     `json:"id"`
	CreatedAt           time.Time `json:"-"`
	Name                string    `json:"name"`
	Type                string    `json:"type,omitempty"`
	IsColor             bool      `json:"is_color,omitempty"`
	IPAddress           string    `json:"ip_address,omitempty"`
	Status              string    `json:"status"`
	SupportedPaperSizes []string  `json:"supported_paper_sizes"`
	Description         string    `json:"description,omitempty"`
	BatteryLeft         Runtime   `json:"battery_left,omitempty"` // -1 means not chargeable
	Version             int32     `json:"version"`
}

func ValidatePrinter(v *validator.Validator, printer *Printer) {
	v.Check(printer.Name != "", "name", "must be provided")
	v.Check(len(printer.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(printer.BatteryLeft != 0, "battery_left", "must be provided")

	v.Check(printer.SupportedPaperSizes != nil, "supported_paper_sizes", "must be provided")
	v.Check(len(printer.SupportedPaperSizes) > 0, "supported_paper_sizes", "must contain at least 1 supported paper size")
	v.Check(validator.Unique(printer.SupportedPaperSizes), "supported_paper_sizes", "must not contain duplicate values")
}

type PrinterModel struct {
	DB *sql.DB
}

func (p PrinterModel) Insert(printer *Printer) error {
	query := `
		INSERT INTO printers (name, type, is_color, ip_address, status, supported_paper_sizes, description, battery_left)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, version`
	args := []interface{}{printer.Name, printer.Type, printer.IsColor, printer.IPAddress, printer.Status, pq.Array(printer.SupportedPaperSizes), printer.Description, printer.BatteryLeft}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return p.DB.QueryRowContext(ctx, query, args...).Scan(&printer.ID, &printer.CreatedAt, &printer.Version)
}

func (p PrinterModel) Get(id int64) (*Printer, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, name, type, is_color, ip_address, status, supported_paper_sizes, description, battery_left, version
		FROM printers
		WHERE id = $1`

	var printer Printer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&printer.ID,
		&printer.CreatedAt,
		&printer.Name,
		&printer.Type,
		&printer.IsColor,
		&printer.IPAddress,
		&printer.Status,
		pq.Array(&printer.SupportedPaperSizes),
		&printer.Description,
		&printer.BatteryLeft,
		&printer.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &printer, nil
}

func (p PrinterModel) Update(printer *Printer) error {
	query := `
		UPDATE printers
		SET name = $1, type = $2, ip_address = $3, status = $4, description = $5, battery_left = $6, version = version + 1
		WHERE id = $7 AND version = $8
		RETURNING version`
	args := []interface{}{
		printer.Name,
		printer.Type,
		printer.IPAddress,
		printer.Status,
		printer.Description,
		printer.BatteryLeft,
		printer.ID,
		printer.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&printer.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (p PrinterModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM printers
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
