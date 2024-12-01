package types

type CallbackQuery struct {
	ID              string                    `json:"id"`                          // Unique identifier for this query
	From            *User                     `json:"from"`                        // Sender
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`           // Optional. Message sent by the bot with the callback button that originated the query
	InlineMessageID string                    `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string                    `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string                    `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   string                    `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}
