package main

import (
	"fmt"
	"github.com/asd/asd/internal/data"
	"net/http"
)

func (app *application) createActorsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FistName     string   `json:"firstName"`
		LastName     string   `json:"lastName"`
		DateOfBirth  int32    `json:"dateOfBirth"`
		MoviesCasted []string `json:"moviesCasted"`
	}
	// if there is error with decoding, we are sending corresponding message
	err := app.readJSON(w, r, &input) //non-nil pointer as the target decode destination
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showActorsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	actor := data.Actor{
		ID:        id,
		FirstName: "David",
		LastName:  "Johnson",
		//DateOfBirth:  0,
		MoviesCasted: []string{"Titanic", "Indian Man"},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"actor": actor}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
