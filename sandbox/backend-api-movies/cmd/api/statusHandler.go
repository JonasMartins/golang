package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := AppStatus{
		Status:      "Avaliable",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		app.logger.Println((err))
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string
	}

	_error := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, statusCode, _error, "error")
}
