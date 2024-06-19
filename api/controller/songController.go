package controller

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"main/internal/stavanFirestore"
	"net/http"

	"github.com/gorilla/mux"
)

// Fetches the song and its details
func GetSong(w http.ResponseWriter, r *http.Request) {
	songId := mux.Vars(r)["songId"]

	if songId == "" {
		log.Default().Println("Throwing 400 Bad Request: songId parameter cannot be empty!")
		//Throw error that playlistTag is not provided.
		http.Error(w, "400 Bad Request: pass songId parameter!", http.StatusBadRequest)
		return
	}

	var song models.Song
	song, err := stavanFirestore.GetSongsData(songId)

	if err != nil {
		log.Default().Println("Throwing 500 Internal Server Error:", err)
		//It happens when the there might be an error in code or the data from the database is not interpreted properly.
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(song)
}
