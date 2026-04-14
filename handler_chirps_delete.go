package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/wet-senker/chirpy2/internal/auth"
	"github.com/wet-senker/chirpy2/internal/database"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}
	rowsDeleted, err := cfg.db.DeleteChirp(
	r.Context(),
	database.DeleteChirpParams{
		ID:     chirpID,
		UserID: userID,
	},
)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}

	if rowsDeleted == 0 {
		respondWithError(w, http.StatusForbidden, "You can't delete this chirp", nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
