package main

import (
	"Printers.imangalizhumash.net/internal/data"
	"Printers.imangalizhumash.net/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPrinterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name                string       `json:"name"`
		Type                string       `json:"type"`
		IsColor             bool         `json:"is_color"`
		Status              string       `json:"status"`
		Description         string       `json:"description"`
		SupportedPaperSizes []string     `json:"supported_paper_sizes"`
		BatteryLeft         data.Runtime `json:"battery_left"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	printer := &data.Printer{
		Name:                input.Name,
		Type:                input.Type,
		IsColor:             input.IsColor,
		Status:              input.Status,
		Description:         input.Description,
		SupportedPaperSizes: input.SupportedPaperSizes,
		BatteryLeft:         input.BatteryLeft,
	}

	v := validator.New()

	if data.ValidatePrinter(v, printer); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showPrinterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// Use the new notFoundResponse() helper.
		app.notFoundResponse(w, r)
		return
	}

	printer := data.Printer{
		ID:                  id,
		CreatedAt:           time.Now(),
		Name:                "Printer 1",
		Type:                "Laser",
		IsColor:             true,
		IPAddress:           "1.1.1.1",
		Status:              "online",
		Description:         "Nice, cheap printer",
		SupportedPaperSizes: []string{"A4", "A3"},
		BatteryLeft:         10,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"printer": printer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
