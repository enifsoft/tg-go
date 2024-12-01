package types

type PaidMediaVideo struct {
	Type  string `json:"type"`  // Type of the paid media, always “video”
	Video *Video `json:"video"` // The video
}
