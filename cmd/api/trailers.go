package main

import (
	"github.com/asd/asd/internal/data"
	"net/http"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint.
// return a JSON response.
func (app *application) createTrailerHandler(w http.ResponseWriter, r *http.Request) {
	//Declare an anonymous struct to hold the information that we expect to be in the
	// HTTP request body (note that the field names and types in the struct are a subset
	// of the Movie struct that we created earlier). This struct will be our *target
	// decode destination*.
	var input struct {
		TrailerName string `json:"trailer_name"`
		Duration    int32  `json:"duration"`
		PremierDate string `json:"premier_date"`
	}
	// if there is error with decoding, we are sending corresponding message
	err := app.readJSON(w, r, &input) //non-nil pointer as the target decode destination
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	trailer := &data.Trailer{
		TrailerName: input.TrailerName,
		Duration:    input.Duration,
		PremierDate: input.PremierDate,
	}

	err = app.models.Trailers.Insert(trailer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	//headers := make(http.Header)
	//headers.Set("Location", fmt.Sprintf("/v1/trailers/%d", trailer.ID))
	// Write a JSON response with a 201 Created status code, the movie data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"trailer": trailer}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
