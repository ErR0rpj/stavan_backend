package stavanFirestore

import (
	"errors"
	"fmt"
	"log"

	config "main/config"
	models "main/internal/models"

	"google.golang.org/api/iterator"
)

// /Fetches all the playlist from firestore
func GetAllPlaylists() ([]models.Playlist, error) {
	fmt.Println("Getting all the playlists from firestore")

	iter := config.CLIENT.Collection("playlists").Documents(config.CTX)
	var playlists []models.Playlist
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var playlist models.Playlist
		doc.DataTo(&playlist)

		playlist, err = models.ValidatePlaylist(playlist)
		if err != nil {
			return nil, err
		} else {
			playlists = append(playlists, playlist)
		}

	}

	return playlists, nil
}

// Gets the songs data (not the likes, shares, etc count) for a particular song
func GetSongsData(songId string) (*models.DynamicBanner, error) {
	fmt.Println("Getting songs data for", songId)

	docSnap, err := config.CLIENT.Collection("songs").Doc(songId).Get(config.CTX)
	if err != nil {
		log.Default().Println("Error getting song data from firestore. Check the songId!", err)
		return nil, errors.New("Error getting song data from firestore. Check the songId!: " + err.Error())
	}

	songMap := docSnap.Data()

	dynamicBanner := models.DynamicBanner{
		Id:         songId,
		BannerType: songMap["category"].(string),
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}

	return &dynamicBanner, nil
}
