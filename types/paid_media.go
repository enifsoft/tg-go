package types

type PaidMedia struct {
	Type     string `json:"type"`               // Type of the paid media, always “preview”
	Width    int32  `json:"width,omitempty"`    // Optional. Media width as defined by the sender
	Height   int32  `json:"height,omitempty"`   // Optional. Media height as defined by the sender
	Duration int32  `json:"duration,omitempty"` // Optional. Duration of the media in seconds as defined by the sender
}
