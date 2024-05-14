package main

import (
	"net/http"

	"github.com/ArrayOfLilly/blog-aggregator/internal/auth"
)

func (apiCfg *apiConfig) handleUsersGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find API key")
		return
	}
	
	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNoContent, "Couldn't find user")
	}
	
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}