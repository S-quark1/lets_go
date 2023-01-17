package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//movies
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMoviesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	//actors
	router.HandlerFunc(http.MethodPost, "/v1/actors", app.createActorsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/actors/:id", app.showActorsHandler)

	// trailers
	router.HandlerFunc(http.MethodGet, "/v1/trailers", app.listTrailersHandler)
	router.HandlerFunc(http.MethodPost, "/v1/trailers", app.createTrailerHandler)

	// users
	//router.HandlerFunc(http.MethodGet, "/v1/trailers", app.listTrailersHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	//tokens
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// Wrap the router with the panic recovery middleware.
	//return app.recoverPanic(app.authenticate(router))
	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
