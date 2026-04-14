package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := cfg.db.Reset(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete user", err)
		return
	}
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
