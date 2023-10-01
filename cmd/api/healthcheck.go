package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// create a map to hold info we want to send
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	// use json.Marshal() to encode JSON
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		// use serverErrorResponse() helper
		app.serverErrorResponse(w, r, err)
	}
}
