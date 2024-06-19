package stavanFirestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	config "main/config"
	models "main/internal/models"

	"google.golang.org/api/iterator"
)

// Fetches all the playlist from firestore
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
		bytes, _ := json.Marshal(doc.Data())
		err = json.Unmarshal(bytes, &playlist)
		if err != nil {
			fmt.Println(err)
		}

		playlist, err = models.ValidatePlaylist(playlist)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)

	}

	return playlists, nil
}

// /Fetches all the playlist from firestore
func GetSongsFromPlaylist(playlistTag string) ([]models.Song, error) {
	fmt.Println("Getting all the playlists from firestore")

	//When getting data from the firestore id = playlistTag.
	doc, err := config.CLIENT.Collection("playlists").Doc(playlistTag).Get(config.CTX)
	if err != nil {
		log.Default().Println("Error getting playlist data from firestore. Check the playlistTag!", err)
		return nil, errors.New("Error getting songs for playlist: " + err.Error())
	}

	var playlist models.Playlist
	bytes, _ := json.Marshal(doc.Data())
	err = json.Unmarshal(bytes, &playlist)
	if err != nil {
		fmt.Println(err)
	}

	playlist, err = models.ValidatePlaylist(playlist)
	if err != nil {
		return nil, err
	}

	songIds := playlist.Songs
	if len(songIds) == 0 {
		return nil, errors.New("No songs found for playlist: " + playlistTag)
	}

	var songs []models.Song
	for _, songId := range songIds {
		song, err := GetSongsData(songId)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

// Gets the songs data (not the likes, shares, etc count) for a particular song
func GetSongsData(songId string) (models.Song, error) {
	fmt.Println("Getting songs data for", songId)

	var song models.Song
	doc, err := config.CLIENT.Collection("songs").Doc(songId).Get(config.CTX)
	if err != nil {
		log.Default().Println("Error getting song data from firestore. Check the songId!", err)
		return song, errors.New("Error getting song data from firestore. Check the songId!: " + err.Error())
	}

	bytes, _ := json.Marshal(doc.Data())
	err = json.Unmarshal(bytes, &song)
	if err != nil {
		fmt.Println(err)
	}

	song, err = models.ValidateSong(song)
	if err != nil {
		return song, err
	}

	return song, nil
}
