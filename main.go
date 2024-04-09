package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Dynamic banner model for backend. It is shown in song page in the app.
type DynamicBanner struct {
	Id         string `json:"id"`
	BannerType string `json:"banner_type"` //Type can be: Story, Advertisement, Playlist, Pachhkhan
	ItemId     string `json:"item_id"`     //This is the id of the element which needs to be fetched. The ID can be NA or blank in case of advertisement
	//Can be: Online (Backend/Firebase), Offline (None).
	FetchFrom string `json:"fetch_from"`
}

// It calculates which banner should be shown for the particular song and then creates the dynamic banner and returns it.
func createDynamicBanner(w http.ResponseWriter, r *http.Request) {

	requstMap := mux.Vars(r)

	fmt.Println("Creating dynamic banner for", requstMap["songId"])

	dynamicBanner := DynamicBanner{
		Id:         "pachchhkhan_bottom_sheet",
		BannerType: "Pachhkhan",
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(dynamicBanner)
}

func handleRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/get-dynamic-banner/{songId}", createDynamicBanner).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	handleRoutes()
}
