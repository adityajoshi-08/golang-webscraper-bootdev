package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusOK, "something went terribly wrong")
}
