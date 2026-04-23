package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/wet-senker/chirpy2/internal/auth"
)

func (cfg *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userPolka, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api Key", err)
		return
	}
	if userPolka != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "No Payment found", err)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	rows, err := cfg.db.UpgradeUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't upgrade user", err)
		return
	}
	if rows == 0 {
		respondWithError(w, http.StatusNotFound, "User not found in database", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
