package models

import (
	"errors"
	"fmt"
)

// Dynamic banner model for backend. It is shown in song page in the app.
type Playlist struct {
	///This can be _id in case of mongo otherwise id in case of firebase
	Id          string `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	PlaylistTag string `json:"playlistTag"`
	//This can be: genre, category, tirthankar, etc.
	PlaylistTagType  string   `json:"playlistTagType"`
	Songs            []string `json:"songs"`
	TodayViews       int      `json:"todayViews"`
	TotalViews       int      `json:"totalViews"`
	Popularity       int      `json:"popularity"`
	Shares           int      `json:"shares"`
	TrendPoints      float64  `json:"trendPoints"`
	LastModifiedTime string   `json:"lastModifiedTime"`
}

// This validates if the playlist object can be formed or not.
// This should be used whenever an playlist object is created.
func ValidatePlaylist(playlist Playlist) (Playlist, error) {
	if playlist.Id == "" || playlist.Title == "" {
		fmt.Println("Playlist has been invalidated: ", playlist.Id, playlist.Title)
		return playlist, errors.New("playlist invalidated")
	} else {
		return playlist, nil
	}
}
