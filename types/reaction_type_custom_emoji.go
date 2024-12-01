package types

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`            // Type of the reaction, always “custom_emoji”
	CustomEmojiID string `json:"custom_emoji_id"` // Custom emoji identifier
}
