package main

import (
	"net/http"

	"github.com/adityajoshi-08/golang-webscraper-bootdev/internal/auth"
	"github.com/adityajoshi-08/golang-webscraper-bootdev/internal/database"
	"github.com/google/uuid"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeySTR, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "missing or invalid API key")
		}
		apiKey, err := uuid.Parse(apiKeySTR)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error parsing API key")
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error fetching user by API key")
			return
		}
		handler(w, r, user)
	}

}
