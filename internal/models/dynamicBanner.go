package models

// Dynamic banner model for backend. It is shown in song page in the app.
type DynamicBanner struct {
	Id         string `json:"id"`
	BannerType string `json:"banner_type"` //Type can be: Story, Advertisement, Playlist, Pachhkhan
	ItemId     string `json:"item_id"`     //This is the id of the element which needs to be fetched. The ID can be NA or blank in case of advertisement
	//Can be: Online (Backend/Firebase), Offline (None).
	FetchFrom string `json:"fetch_from"`
}
