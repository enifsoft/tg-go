package types

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`   // Optional. Title text of the business intro
	Message string   `json:"message,omitempty"` // Optional. Message text of the business intro
	Sticker *Sticker `json:"sticker,omitempty"` // Optional. Sticker of the business intro
}
