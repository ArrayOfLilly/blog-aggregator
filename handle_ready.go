package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	status := http.StatusText(http.StatusOK)
	respondWithJSON(w, http.StatusOK, map[string]string{"Status": status})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}