package models

import "fmt"

// Dynamic banner model for backend. It is shown in song page in the app.
type Playlist struct {
	///This can be _id in case of mongo otherwise id in case of firebase
	Id          string `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	PlaylistTag string `json:"playlist_tag"`
	///This can be: genre, category, tirthankar, etc.
	PlaylistTagType string   `json:"playlist_tag_type"`
	Songs           []string `json:"songs"`
}

//This validates if the playlist object can be formed or not.
//This should be used whenever an playlist object is created.
func ValidatePlaylist(playlist *Playlist) *Playlist {
	if playlist.Id == "" || playlist.Title == "" {
		fmt.Println("Playlist has been invalidated: ", playlist.Id, playlist.Title)
		return nil
	} else {
		return playlist
	}
}
