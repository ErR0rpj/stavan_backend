package controller

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"main/internal/stavanFirestore"
	"net/http"

	"github.com/gorilla/mux"
)

// Fetches all the playlist from the database and sends back the response in json.
func GetAllplaylist(w http.ResponseWriter, r *http.Request) {
	var playlists []models.Playlist
	playlists, err := stavanFirestore.GetAllPlaylists()

	if err != nil {
		log.Default().Println("Throwing 500 Internal Server Error:", err)
		//It happens when the there might be an error in code or the data from the database is not interpreted properly.
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	} else if len(playlists) == 0 {
		log.Default().Println("Throwing 501 Not Implemented Error: Playlists list is nil, error might be in fetching firestore. URL: " + r.URL.String())
		//Throws 501 as the playlist list was empty. This might be due to playlist list actually being empty
		//or the query to database is wrong.
		http.Error(w, "501 Not Implemented Error", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(playlists)
}

// Fetches all the playlist from the database and sends back the response in json.
func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	playlistTag := mux.Vars(r)["playlistTag"]

	if playlistTag == "" {
		log.Default().Println("Throwing 400 Bad Request: playlistTag parameter cannot be empty!")
		//Throw error that playlistTag is not provided.
		http.Error(w, "400 Bad Request: pass playlistTag parameter!", http.StatusBadRequest)
		return
	}

	var songs []models.Song
	songs, err := stavanFirestore.GetSongsFromPlaylist(playlistTag)

	if err != nil {
		log.Default().Println("Throwing 500 Internal Server Error:", err)
		//It happens when the there might be an error in code or the data from the database is not interpreted properly.
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	} else if len(songs) == 0 {
		log.Default().Println("Throwing 501 Not Implemented Error: Songs list is nil, error might be in fetching firestore: " + r.URL.String())
		//Throws 501 as the playlist list was empty. This might be due to playlist list actually being empty
		//or the query to database is wrong.
		http.Error(w, "501 Not Implemented Error", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(songs)
}
