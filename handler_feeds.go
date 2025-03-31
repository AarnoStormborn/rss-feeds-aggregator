package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AarnoStormborn/go-RSS-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, feed)
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error fetching feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, feeds)

}

// func (apiCfg *apiConfig) handlerGetFeedById(w http.ResponseWriter, r *http.Request, user database.User) {

// 	type parameters struct {
// 		ID     uuid.UUID `json:"id"`
// 		UserID uuid.UUID `json:"user_id"`
// 	}

// 	args := parameters{ID: }

// 	feed, err := apiCfg.DB.GetFeed(r.Context())

// 	respondWithJSON(w, 200, user)
// }
