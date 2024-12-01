package types

type UserChatBoosts struct {
	Boosts []*ChatBoost `json:"boosts"` // The list of boosts added to the chat by the user
}
