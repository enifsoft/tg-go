package types

type PaidMediaPhoto struct {
	Type  string       `json:"type"`  // Type of the paid media, always “photo”
	Photo []*PhotoSize `json:"photo"` // The photo
}
