package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/william-cirico/rest-api-golang/errors"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
	return
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

type RootHandler func(w http.ResponseWriter, r *http.Request) error

func (ch RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := ch(w, r); err != nil {
		httpError, ok := err.(*errors.HTTPError)
		if ok {
			respondWithError(w, httpError.StatusCode, httpError.Message)
		} else {
			fmt.Println(httpError.Cause)
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}
}
