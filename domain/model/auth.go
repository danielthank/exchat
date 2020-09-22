package model

type Profile struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureURL  string `json:"pictureUrl"`
	AccessToken string `json:"-"`
}
