package models

import (
	"errors"
	"fmt"
)

// Song model
type Song struct {
	//This can be _id in case of mongo otherwise id in case of firebase.
	//This was called code in flutter models
	Id string `json:"id"`
	//This is the name of the film or the album from which song is based on
	Album string `json:"album"`
	//This can be: Latest, Bhakti, Stotra, Stavan, Song, etc
	Category string `json:"category"`
	//This can be: Tapasya, Diksha, Paryushan, Girnar, Palitana, etc
	Genre string `json:"genre"`
	//This can be: Mahavir, Parshwanath, Neminath, Nakoda Bheru, etc
	Tirthankar string `json:"tirthankar"`
	Lyrics1    string `json:"lyrics1"`
	Lyrics2    string `json:"lyrics2"`
	Lyrics3    string `json:"lyrics3"`
	//This is the language in which the song is sung
	Language string `json:"language"`
	//This is the song on which the stavan is based on
	OriginalSong string `json:"originalSong"`
	Likes        int    `json:"likes"`
	Popularity   int    `json:"popularity"`
	Shares       int    `json:"shares"`
	//This is the production house who own the rights of the stavan
	Production       string   `json:"production"`
	SearchKeywords   string   `json:"searchKeywords"`
	SingerName       string   `json:"singerName"`
	SongNameEnglish  string   `json:"songNameEnglish"`
	SongNameHindi    string   `json:"songNameHindi"`
	TodayViews       string   `json:"todayViews"`
	TotalViews       string   `json:"totalViews"`
	TrendPoints      float64  `json:"trendPoints"`
	YoutubeLink      string   `json:"youtubeLink"`
	StoryLinks       []string `story:"storyLinks"`
	LastModifiedTime int      `json:"lastModifiedTime"`
}

// This validates if the song object can be formed or not.
// This should be used whenever a song object is created.
func ValidateSong(song Song) (Song, error) {
	if song.Id == "" || song.SongNameEnglish == "" {
		fmt.Println("Song has been invalidated: ", song.Id, song.SongNameEnglish)
		return song, errors.New("song invalidated")
	} else {
		return song, nil
	}
}
