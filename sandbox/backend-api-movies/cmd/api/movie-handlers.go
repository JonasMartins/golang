package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type jesonResp struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

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

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type MoviePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	err := app.models.DB.DeleteMovie(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	ok := jesonResp{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

	var payload MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var movie models.Movie

	if payload.ID != "" {
		m, _ := app.models.DB.Get(payload.ID)
		movie = *m
		movie.UpdatedAt = time.Now()
	}

	movie.ID = payload.ID
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2000-01-01", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	var rating, _ = strconv.ParseFloat(payload.Rating, 32)
	movie.Rating = float32(rating)
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	if movie.ID == "" {
		err = app.models.DB.InertMovie(movie)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateMovie(movie)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}
	ok := jesonResp{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}
func (app *application) insertMovie(w http.ResponseWriter, r *http.Request)  {}
func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {}
