package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var client *firestore.Client
var ctx context.Context

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
	songId := requstMap["songId"]
	fmt.Println("Creating dynamic banner for", songId)
	dynamicBanner := getSongsDataFromFirebase(songId)

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(dynamicBanner)
}

//http://192.168.1.4:8082/get-dynamic-banner/TMB

// Starts the server and starts APIs
func handleRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/get-dynamic-banner/{songId}", createDynamicBanner).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

// Initialzies the firebase with service account details
func initializeServiceAccountForFirebase() *firestore.Client {
	fmt.Println("Initialize service account for Firebase.")

	sa := option.WithCredentialsFile("service-account-credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Service account initialized.")
	fmt.Println("")

	return client
}

func getSongsDataFromFirebase(songId string) DynamicBanner {
	fmt.Println("Getting songs data for", songId)

	docSnap, err := client.Collection("songs").Doc(songId).Get(ctx)
	if err != nil {
		log.Fatalln("Error getting song data from firestore", err)
	}

	songMap := docSnap.Data()

	dynamicBanner := DynamicBanner{
		Id:         songId,
		BannerType: songMap["category"].(string),
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}

	return dynamicBanner
}

func main() {
	//Get firestore client
	ctx = context.Background()
	client = initializeServiceAccountForFirebase()
	defer client.Close()

	handleRoutes()

	defer client.Close()
}
