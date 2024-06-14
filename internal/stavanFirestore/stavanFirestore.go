package stavanFirestore

import (
	"errors"
	"fmt"
	"log"

	config "main/config"
	dynamicBanner "main/internal/models"
)

// Gets the songs data (not the likes, shares, etc count) for a particular song
func GetSongsDataFromFirebase(songId string) (*dynamicBanner.DynamicBanner, error) {
	fmt.Println("Getting songs data for", songId)

	docSnap, err := config.CLIENT.Collection("songs").Doc(songId).Get(config.CTX)
	if err != nil {
		log.Default().Println("Error getting song data from firestore. Check the songId!", err)
		return nil, errors.New("Error getting song data from firestore. Check the songId!: " + err.Error())
	}

	songMap := docSnap.Data()

	dynamicBanner := dynamicBanner.DynamicBanner{
		Id:         songId,
		BannerType: songMap["category"].(string),
		ItemId:     "pachhkhan",
		FetchFrom:  "None",
	}

	return &dynamicBanner, nil
}
