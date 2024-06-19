package admin

import (
	"encoding/json"
	"fmt"
	adminConfig "main/admin/config"
	adminModels "main/admin/models"
	"main/config"
	"main/internal/models"

	"google.golang.org/api/iterator"
)

// Reads the data from json of flutter firestore converts the variable names and stores them to web firestore
func TransferDataFromFlutterFirestoreToWebFirestore() error {
	fmt.Println("Transferring data from flutter firestore to web firestore")

	//Closing previous firebase connection and making new connection with flutter firebase
	config.CLIENT.Close()
	config.CLIENT = adminConfig.InitializeServiceAccountForFlutterFirebase()
	fluterSongs, err := ReadFlutterFirestoreSongs()

	if err != nil {
		//Closing flutter firebase connection and making connection again with web firebase
		config.CLIENT.Close()
		config.CLIENT = config.InitializeServiceAccountForFirebase()
		return err
	}

	//Closing flutter firebase connection and making connection again with web firebase
	config.CLIENT.Close()
	config.CLIENT = config.InitializeServiceAccountForFirebase()

	webSongs, err := ConvertFlutterSongsToWebSongs(fluterSongs)
	if err != nil {
		return err
	}

	err = WriteMultipleSongs(webSongs)
	if err != nil {
		return err
	}

	fmt.Println("Data transfer from flutter firestore to web firestore completed successfully")
	return nil
}

// Converts all the flutter song to web songs models
func ConvertFlutterSongsToWebSongs(flutterSongs []adminModels.SongFlutter) ([]models.Song, error) {
	fmt.Println("Converting flutter songs to web songs")

	var webSongs []models.Song
	for _, song := range flutterSongs {
		var webSong models.Song
		webSong.Id = song.Code
		webSong.Album = song.Album
		webSong.Category = song.Category
		webSong.Genre = song.Genre
		webSong.Tirthankar = song.Tirthankar
		webSong.Lyrics1 = song.Lyrics
		webSong.Lyrics2 = song.GujaratiLyrics
		webSong.Lyrics3 = song.EnglishLyrics
		webSong.Language = song.Language
		webSong.OriginalSong = song.OriginalSong
		webSong.Likes = 0
		webSong.Popularity = 0
		webSong.Shares = 0
		webSong.Production = song.Production
		webSong.SearchKeywords = song.SearchKeywords
		webSong.SingerName = song.Singer
		webSong.SongNameEnglish = song.SongNameEnglish
		webSong.SongNameHindi = song.SongNameHindi
		webSong.TodayViews = 0
		webSong.TotalViews = 0
		webSong.TrendPoints = 0.0
		webSong.YoutubeLink = song.YoutubeLink
		webSong.LastModifiedTime = song.LastModifiedTime

		webSong, err := models.ValidateSong(webSong)
		if err != nil {
			return nil, err
		}
		webSongs = append(webSongs, webSong)
	}
	return webSongs, nil
}

// Reads the songs from the flutter firebase and returns the list of all the songs.
func ReadFlutterFirestoreSongs() ([]adminModels.SongFlutter, error) {
	fmt.Println("Reading flutter firestore songs")

	iter := config.CLIENT.Collection("songs").Documents(config.CTX)
	var songs []adminModels.SongFlutter
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var song adminModels.SongFlutter
		bytes, _ := json.Marshal(doc.Data())
		err = json.Unmarshal(bytes, &song)
		if err != nil {
			fmt.Println(err)
		}

		song, err = adminModels.ValidateSongFlutter(song)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}
