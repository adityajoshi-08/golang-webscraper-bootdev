package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adityajoshi-08/golang-webscraper-bootdev/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID  string `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID: uuid.MustParse(params.FeedID),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create feed follow")
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseToFeedFollow(feed))
}
