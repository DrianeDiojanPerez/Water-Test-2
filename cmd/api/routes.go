// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//create a new httprouter router instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/waterbill",app.requirePermission("waters:read",app.waterbill_listHandler))

	router.HandlerFunc(http.MethodPost, "/v1/waterbill", app.requirePermission("waters:write",app.createwaterbill_listHandler))
	router.HandlerFunc(http.MethodGet, "/v1/waterbill/:id", app.requirePermission("waters:read",app.showwaterbill_listHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/waterbill/:id", app.requirePermission("waters:write",app.updatewaterbill_listHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/waterbill/:id", app.requirePermission("waters:write",app.deletewaterbill_listItemHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	//we wrap router with recoverpanic will call router if everthing is okay
	//then we pass to the rate limit and the process the actual request
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
