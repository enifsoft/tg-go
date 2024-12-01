package types

type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`                 // Type of the background, always “wallpaper”
	Document         *Document `json:"document"`             // Document with the wallpaper
	DarkThemeDimming int32     `json:"dark_theme_dimming"`   // Dimming of the background in dark themes, as a percentage; 0-100
	IsBlurred        bool      `json:"is_blurred,omitempty"` // Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsMoving         bool      `json:"is_moving,omitempty"`  // Optional. True, if the background moves slightly when the device is tilted
}
