package types

type ForumTopicClosed struct {
	Name              string `json:"name,omitempty"`                 // Optional. New name of the topic, if it was edited
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"` // Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
}
