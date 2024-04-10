package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// Initialzies the firebase with service account details
func initializeServiceAccountForFirebase() *firestore.Client {
	fmt.Println("Initializing service account for Firebase.")

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

	return client
}

// Starts the server and starts APIs
func handleRoutes() {
	fmt.Println("Starting the server on port 8082...")
	fmt.Println("")

	router := mux.NewRouter()
	router.HandleFunc("/get-dynamic-banner", createDynamicBanner).Methods("GET")

	//This creates a server at the port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

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
	var dynamicBanner *DynamicBanner
	dynamicBanner, err = getSongsDataFromFirebase(songId)
	if err != nil {
		log.Default().Println("Throwing 400 Bad Request:", err)
		//throw error that isLiked should either be true/false
		http.Error(w, "400 Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	} else if dynamicBanner == nil {
		log.Default().Println("Throwing 500 Internal Server Error: Dynamic banner is nil, error might be in fetching firestore. URL: " + r.URL.String())
		//throw error that isLiked should either be true/false
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/")
	json.NewEncoder(w).Encode(dynamicBanner)
}

//http://192.168.1.4:8082/get-dynamic-banner/TMB&false

func getSongsDataFromFirebase(songId string) (*DynamicBanner, error) {
	fmt.Println("Getting songs data for", songId)

	docSnap, err := client.Collection("songs").Doc(songId).Get(ctx)
	if err != nil {
		log.Default().Println("Error getting song data from firestore. Check the songId!", err)
		return nil, errors.New("Error getting song data from firestore. Check the songId!: " + err.Error())
	}

	songMap := docSnap.Data()

	dynamicBanner := DynamicBanner{
		Id:         songId,
		BannerType: songMap["category"].(string),
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}

	return &dynamicBanner, nil
}

func main() {
	//Get firestore client
	ctx = context.Background()
	client = initializeServiceAccountForFirebase()
	defer client.Close()

	handleRoutes()

	defer client.Close()
}
