package types

type ChatBoostSourcePremium struct {
	Source string `json:"source"` // Source of the boost, always “premium”
	User   *User  `json:"user"`   // User that boosted the chat
}
