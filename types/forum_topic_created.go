package types

type ForumTopicCreated struct {
	Name              string `json:"name"`                           // Name of the topic
	IconColor         int32  `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}
