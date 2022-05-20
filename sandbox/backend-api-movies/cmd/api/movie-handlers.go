package main

import (
	"backend/models"
	"errors"
	"net/http"
	"time"

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

	movie := models.Movie{
		ID:          "1",
		Title:       "Test",
		Description: "test",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 0, 0, 0, 0, time.Local),
		Runtime:     90,
		Rating:      5.5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
