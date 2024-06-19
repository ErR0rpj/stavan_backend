package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/stavanFirestore"
	"net/http"
	"strconv"
)

// It calculates which banner should be shown for the particular song and then creates the dynamic banner and returns it.
func CreateDynamicBanner(w http.ResponseWriter, r *http.Request) {
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
