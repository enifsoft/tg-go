package types

type BackgroundType struct {
	Type             string          `json:"type"`               // Type of the background, always “fill”
	Fill             *BackgroundFill `json:"fill"`               // The background fill
	DarkThemeDimming int32           `json:"dark_theme_dimming"` // Dimming of the background in dark themes, as a percentage; 0-100
}
