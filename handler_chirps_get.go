package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/carterdr/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}
	author_id, err := strconv.Atoi(r.URL.Query().Get("author_id"))
	if err != nil {
		author_id = -1
	}
	chirps := []database.Chirp{}
	for _, dbChirp := range dbChirps {
		if dbChirp.AuthorID == author_id || author_id == -1 {
			chirps = append(chirps, database.Chirp{
				ID:   dbChirp.ID,
				Body: dbChirp.Body,
			})
		}

	}
	sortOrder := r.URL.Query().Get("sort")
	if sortOrder == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID > chirps[j].ID
		})
	} else {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID < chirps[j].ID
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpRetrieve(w http.ResponseWriter, r *http.Request) {
	chirpID, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID")
		return
	}
	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, database.Chirp{
		ID:   dbChirp.ID,
		Body: dbChirp.Body,
	})
}
