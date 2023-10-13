package data

import "time"

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
