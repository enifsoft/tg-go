package types

type BackgroundFillFreeformGradient struct {
	Type   string  `json:"type"`   // Type of the background fill, always “freeform_gradient”
	Colors []int32 `json:"colors"` // A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
}
