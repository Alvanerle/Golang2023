package main

import (
	"Printers.imangalizhumash.net/internal/data"
	"Printers.imangalizhumash.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createPrinterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name                string       `json:"name"`
		Type                string       `json:"type"`
		IsColor             bool         `json:"is_color"`
		IPAddress           string       `json:"ip_address"`
		Status              string       `json:"status"`
		SupportedPaperSizes []string     `json:"supported_paper_sizes"`
		Description         string       `json:"description"`
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
		IPAddress:           input.IPAddress,
		Status:              input.Status,
		SupportedPaperSizes: input.SupportedPaperSizes,
		Description:         input.Description,
		BatteryLeft:         input.BatteryLeft,
	}

	v := validator.New()
	if data.ValidatePrinter(v, printer); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Printers.Insert(printer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/printers/%d", printer.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"printer": printer}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPrinterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	printer, err := app.models.Printers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"printer": printer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePrinterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	printer, err := app.models.Printers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name        *string       `json:"name"`
		Type        *string       `json:"type"`
		IpAddress   *string       `json:"ip_address"`
		Status      *string       `json:"status"`
		Description *string       `json:"description"`
		BatteryLeft *data.Runtime `json:"battery_left"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		printer.Name = *input.Name
	}
	if input.Type != nil {
		printer.Type = *input.Type
	}
	if input.IpAddress != nil {
		printer.IPAddress = *input.IpAddress
	}
	if input.Status != nil {
		printer.Status = *input.Status
	}
	if input.Description != nil {
		printer.Description = *input.Description
	}
	if input.BatteryLeft != nil {
		printer.BatteryLeft = *input.BatteryLeft
	}

	v := validator.New()
	if data.ValidatePrinter(v, printer); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Printers.Update(printer)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"printer": printer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) deletePrinterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Printers.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "printer successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPrintersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name                string
		Type                string
		SupportedPaperSizes []string
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Type = app.readString(qs, "type", "")
	input.SupportedPaperSizes = app.readCSV(qs, "supported_paper_sizes", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"name", "type", "status", "-name", "-type", "-status"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
