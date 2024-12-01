package types

type BackgroundFillSolid struct {
	Type  string `json:"type"`  // Type of the background fill, always “solid”
	Color int32  `json:"color"` // The color of the background fill in the RGB24 format
}
