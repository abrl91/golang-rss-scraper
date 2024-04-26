package main

import (
	"net/http"

	"github.com/abrl91/golang-rss-scraper/internal/auth"
	"github.com/abrl91/golang-rss-scraper/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	}
}
