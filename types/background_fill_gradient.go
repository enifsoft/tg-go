package types

type BackgroundFillGradient struct {
	Type          string `json:"type"`           // Type of the background fill, always “gradient”
	TopColor      int32  `json:"top_color"`      // Top color of the gradient in the RGB24 format
	BottomColor   int32  `json:"bottom_color"`   // Bottom color of the gradient in the RGB24 format
	RotationAngle int32  `json:"rotation_angle"` // Clockwise rotation angle of the background fill in degrees; 0-359
}
