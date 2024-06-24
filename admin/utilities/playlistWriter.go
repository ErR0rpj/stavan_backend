package utilities

import (
	"fmt"
	"main/config"
	"main/internal/models"
	"main/internal/stavanFirestore"
	"strings"
)

// This function creates the playlist and add songs in it too.
// It creates the new playlists in the firebase and writes them as batches
// Whatever playlist you want to create new, add them inside the define playlist method.
func WritePlaylistsWithSongs() error {
	//Whatever playlist you want to create new, add them inside the define playlist method.
	playlistList := definePlaylists()
	songsList, err := stavanFirestore.GetAllSongs()
	if err != nil {
		return err
	}

	fmt.Println("Writing playlists with songs...")
	batch := config.CLIENT.BulkWriter(config.CTX)

	for _, playlist := range playlistList {
		getPlaylistSongs(&playlist, songsList)

		sfRef := config.CLIENT.Collection("playlists").Doc(playlist.Id)
		//Writing each playlist in the batch
		_, err := batch.Create(sfRef, playlist)
		if err != nil {
			return err
		}
	}

	//Commits the batch and writes all the changes
	batch.End()
	return nil
}

// Creates predefined playlists model.
func definePlaylists() []models.Playlist {
	var playlistList []models.Playlist
	var playlist models.Playlist

	//Latest playlist
	playlist.Id = "latest"
	playlist.Title = "Latest Releases"
	playlist.Subtitle = "New song lyrics"
	playlist.PlaylistTag = "latest"
	playlist.PlaylistTagType = "genre"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Bhakti
	playlist.Id = "bhakti"
	playlist.Title = "Bhakti Special"
	playlist.Subtitle = "All time Favourite Bhakti"
	playlist.PlaylistTag = "bhakti"
	playlist.PlaylistTagType = "category"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//paryushan
	playlist.Id = "paryushan"
	playlist.Title = "Paryushan Stavans"
	playlist.Subtitle = "Paryushan Mahaparv Playlist"
	playlist.PlaylistTag = "paryushan"
	playlist.PlaylistTagType = "genre"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Tapasya
	playlist.Id = "tapasya"
	playlist.Title = "Tapasya Geet"
	playlist.Subtitle = "Varitap parna, Navtap & others"
	playlist.PlaylistTag = "tapasya"
	playlist.PlaylistTagType = "genre"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Diksha
	playlist.Id = "diksha"
	playlist.Title = "Diksha Stavans"
	playlist.Subtitle = "Diksha playlist"
	playlist.PlaylistTag = "diksha"
	playlist.PlaylistTagType = "genre"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Vicky parekh
	playlist.Id = "vicky"
	playlist.Title = "Vicky Parekh Hits"
	playlist.Subtitle = "Vicky Parekh's best songs"
	playlist.PlaylistTag = "vicky"
	playlist.PlaylistTagType = "singer"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Rishabh Sambhav Jain
	playlist.Id = "rsj"
	playlist.Title = "Rishabh Sambhav Jain"
	playlist.Subtitle = "RSJ's best songs"
	playlist.PlaylistTag = "rsj"
	playlist.PlaylistTagType = "singer"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Stotra
	playlist.Id = "stotra"
	playlist.Title = "Stotra"
	playlist.Subtitle = "Famouns Stotras"
	playlist.PlaylistTag = "stotra"
	playlist.PlaylistTagType = "category"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//neminath and girnar
	playlist.Id = "neminath"
	playlist.Title = "Neminath and Girnar"
	playlist.Subtitle = "Neminath and Girnar Bhajans"
	playlist.PlaylistTag = "neminath"
	playlist.PlaylistTagType = "tirthankar"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//parshwanath
	playlist.Id = "parshwanath"
	playlist.Title = "Parshwanath Swami"
	playlist.Subtitle = "Parasnath Bhajans"
	playlist.PlaylistTag = "parshwanath"
	playlist.PlaylistTagType = "tirthankar"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Mahaveer
	playlist.Id = "mahavir"
	playlist.Title = "Mahaveer Swami"
	playlist.Subtitle = "Mahaveer Swami Bhajans"
	playlist.PlaylistTag = "mahavir"
	playlist.PlaylistTagType = "tirthankar"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Adinath
	playlist.Id = "adinath"
	playlist.Title = "Adinath Swami"
	playlist.Subtitle = "Rishabh dev Bhajans"
	playlist.PlaylistTag = "adinath"
	playlist.PlaylistTagType = "tirthankar"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	//Nakoda
	playlist.Id = "nakoda"
	playlist.Title = "Nakoda Bheruji"
	playlist.Subtitle = "Nakoda Bhairav Bhajans"
	playlist.PlaylistTag = "nakoda"
	playlist.PlaylistTagType = "tirthankar"
	playlist.TotalViews = 0
	playlist.TodayViews = 0
	playlist.Popularity = 0
	playlist.Shares = 0
	playlist.TrendPoints = 0.0
	playlistList = append(playlistList, playlist)

	return playlistList
}

// Scans all the songs from the firestore and adds them to the current playlist if conditoins satisfy
// for the list.
func getPlaylistSongs(playlist *models.Playlist, songsList []models.Song) {
	playlistTag := playlist.PlaylistTag
	playlistTagType := playlist.PlaylistTagType
	for _, song := range songsList {
		switch playlistTagType {
		case "genre":
			if strings.Contains(strings.ToLower(song.Genre), strings.ToLower(playlistTag)) {
				playlist.Songs = append(playlist.Songs, song.Id)
			}
		case "category":
			if strings.Contains(strings.ToLower(song.Category), strings.ToLower(playlistTag)) {
				playlist.Songs = append(playlist.Songs, song.Id)
			}
		case "singer":
			if strings.Contains(strings.ToLower(song.SingerName), strings.ToLower(playlistTag)) {
				playlist.Songs = append(playlist.Songs, song.Id)
			}
		case "tirthankar":
			if strings.Contains(strings.ToLower(song.Tirthankar), strings.ToLower(playlistTag)) {
				playlist.Songs = append(playlist.Songs, song.Id)
			}
		}
	}

}
