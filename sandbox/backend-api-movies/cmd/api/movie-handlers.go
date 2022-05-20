package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id := params.ByName("id")
	if len(id) == 0 {
		app.logger.Print(errors.New("invalid id"))
		app.errorJSON(w, errors.New("invalid id"))

		return
	}

	movie, err := app.models.DB.Get(id)

	if err != nil {
		return
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
