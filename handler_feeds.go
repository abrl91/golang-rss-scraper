package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abrl91/golang-rss-scraper/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newFeed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Error creating feed")
		return
	}

	responseWithJson(w, http.StatusCreated, databaseFeedToFeed(newFeed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Error getting feeds")
		return
	}

	responseWithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
