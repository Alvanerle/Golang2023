package main

import (
	"Printers.imangalizhumash.net/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPrinterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new Printer")
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
