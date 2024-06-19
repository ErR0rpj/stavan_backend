package adminModels

import (
	"errors"
	"fmt"
)

// Song model
type SongFlutter struct {
	Album            string  `json:"album"`
	Category         string  `json:"category"`
	Code             string  `json:"code"`
	EnglishLyrics    string  `json:"englishLyrics"`
	Genre            string  `json:"genre"`
	GujaratiLyrics   string  `json:"gujaratiLyrics"`
	Language         string  `json:"language"`
	LastModifiedTime string  `json:"lastModifiedTime"`
	Likes            int     `json:"likes"`
	Lyrics           string  `json:"lyrics"`
	OriginalSong     string  `json:"originalSong"`
	Popularity       int     `json:"popularity"`
	Production       string  `json:"production"`
	SearchKeywords   string  `json:"searchKeywords"`
	Shares           int     `json:"shares"`
	Singer           string  `json:"singer"`
	SongNameEnglish  string  `json:"songNameEnglish"`
	SongNameHindi    string  `json:"songNameHindi"`
	Tirthankar       string  `json:"tirthankar"`
	TodayClicks      int     `json:"todayClicks"`
	TotalClicks      int     `json:"totalClicks"`
	TrendPoints      float64 `json:"trendPoints"`
	YoutubeLink      string  `json:"youtubeLink"`
}

// This validates if the song object can be formed or not.
// This should be used whenever a song object is created.
func ValidateSongFlutter(song SongFlutter) (SongFlutter, error) {
	if song.Code == "" || song.SongNameEnglish == "" {
		fmt.Println("Song has been invalidated: ", song.Code, song.SongNameEnglish)
		return song, errors.New("song invalidated")
	} else {
		return song, nil
	}
}
