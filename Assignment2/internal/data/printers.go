package data

import (
	"Printers.imangalizhumash.net/internal/validator"
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
}

func ValidatePrinter(v *validator.Validator, printer *Printer) {
	v.Check(printer.Name != "", "name", "must be provided")
	v.Check(len(printer.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(printer.BatteryLeft != 0, "battery_left", "must be provided")

	v.Check(printer.SupportedPaperSizes != nil, "supported_paper_sizes", "must be provided")
	v.Check(len(printer.SupportedPaperSizes) > 0, "supported_paper_sizes", "must contain at least 1 supported paper size")
	v.Check(validator.Unique(printer.SupportedPaperSizes), "supported_paper_sizes", "must not contain duplicate values")
}
