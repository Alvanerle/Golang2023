package main

import (
	"fmt"
	"net/http"
)

func (app *application) createPrinterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new Printer")
}

func (app *application) showPrinterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of Printer %d\n", id)
}
