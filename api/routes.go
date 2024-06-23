package api

import (
	"fmt"
	"log"
	"main/api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// Starts the server and starts APIs
func HandleRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/playlists", controller.GetAllplaylist).Methods("GET")
	router.HandleFunc("/playlists/{playlistTag}", controller.GetPlaylistSongs).Methods("GET")
	router.HandleFunc("/songs/{songId}", controller.GetSong).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}
