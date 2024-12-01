package types

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // Source of the boost, always “gift_code”
	User   *User  `json:"user"`   // User for which the gift code was created
}
