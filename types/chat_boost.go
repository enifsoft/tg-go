package types

type ChatBoost struct {
	BoostID        string           `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int32            `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int32            `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	Source         *ChatBoostSource `json:"source"`          // Source of the added boost
}
