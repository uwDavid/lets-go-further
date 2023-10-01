package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// parameters are stored in req context as well
	params := httprouter.ParamsFromContext(r.Context())

	// use ByName() to get value of param from the slice
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Define an envelope type
type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// use json.MarshalIndent to add whitespace for terminal
	// but there's a performance hit
	// js, err := json.Marshal(data)
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	// append newline for easier viewing in terminal
	js = append(js, '\n')

	// loop through header map => add each to http.ResponseWriter header map
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
