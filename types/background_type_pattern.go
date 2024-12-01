package types

type BackgroundTypePattern struct {
	Type       string          `json:"type"`                  // Type of the background, always “pattern”
	Document   *Document       `json:"document"`              // Document with the pattern
	Fill       *BackgroundFill `json:"fill"`                  // The background fill that is combined with the pattern
	Intensity  int32           `json:"intensity"`             // Intensity of the pattern when it is shown above the filled background; 0-100
	IsInverted bool            `json:"is_inverted,omitempty"` // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsMoving   bool            `json:"is_moving,omitempty"`   // Optional. True, if the background moves slightly when the device is tilted
}
