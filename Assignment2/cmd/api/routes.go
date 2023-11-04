package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/printers", app.listPrintersHandler)
	router.HandlerFunc(http.MethodGet, "/v1/printers/:id", app.showPrinterHandler)
	router.HandlerFunc(http.MethodPost, "/v1/printers", app.createPrinterHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/printers/:id", app.updatePrinterHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/printers/:id", app.deletePrinterHandler)

	return app.recoverPanic(app.rateLimit(router))
}
