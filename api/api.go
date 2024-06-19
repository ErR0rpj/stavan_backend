package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "main/internal/models"
	stavanFirestore "main/internal/stavanFirestore"

	"github.com/gorilla/mux"
)

// Starts the server and starts APIs
func HandleRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/playlists", getAllplaylist).Methods("GET")
	router.HandleFunc("/playlists/{playlistTag}", getPlaylistSongs).Methods("GET")
	// router.HandleFunc("/songs", getTrendingSongs).Methods("GET")
	router.HandleFunc("/songs/{songId}", getSong).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

// Fetches the song and its details
func getSong(w http.ResponseWriter, r *http.Request) {
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

// Fetches all the playlist from the database and sends back the response in json.
func getPlaylistSongs(w http.ResponseWriter, r *http.Request) {
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

// Fetches all the playlist from the database and sends back the response in json.
func getAllplaylist(w http.ResponseWriter, r *http.Request) {
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

// http://192.168.1.4:8082/get-dynamic-banner/TMB&false
// It calculates which banner should be shown for the particular song and then creates the dynamic banner and returns it.
func createDynamicBanner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating dynamic banner.")

	requestMap := r.URL.Query()
	var songId string
	var isLiked bool
	var err error

	has := requestMap.Has("songId")
	if has {
		songId = requestMap.Get("songId")
	} else {
		log.Default().Println("Throwing 400 Bad Request: songdId parameter cannot be empty!")
		//Throw error that song id is not provided.
		http.Error(w, "400 Bad Request: songdId parameter cannot be empty!", http.StatusBadRequest)
		return
	}

	has = requestMap.Has("isLiked")
	if has {

		isLiked, err = strconv.ParseBool(requestMap.Get("isLiked"))
		if err != nil {
			log.Default().Println("Throwing 400 Bad Request: isLiked should either be true/false")
			//throw error that isLiked should either be true/false
			http.Error(w, "400 Bad Request: isLiked parameter can either be true/false!", http.StatusBadRequest)
			return
		}
	}

	fmt.Println("Creating dynamic banner for", songId, isLiked)
	var dynamicBanner *models.DynamicBanner
	var song models.Song
	song, err = stavanFirestore.GetSongsData(songId)

	dynamicBanner = &models.DynamicBanner{
		Id:         songId,
		BannerType: song.Category,
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}
	if err != nil {
		log.Default().Println("Throwing 400 Bad Request:", err)
		//throw error that isLiked should either be true/false
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(dynamicBanner)
}
