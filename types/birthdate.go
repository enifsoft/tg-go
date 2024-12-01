package types

type Birthdate struct {
	Day   int32 `json:"day"`            // Day of the user's birth; 1-31
	Month int32 `json:"month"`          // Month of the user's birth; 1-12
	Year  int32 `json:"year,omitempty"` // Optional. Year of the user's birth
}
