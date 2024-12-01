package types

type BusinessConnection struct {
	ID         string `json:"id"`           // Unique identifier of the business connection
	User       *User  `json:"user"`         // Business account user that created the business connection
	UserChatID int64  `json:"user_chat_id"` // Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Date       int32  `json:"date"`         // Date the connection was established in Unix time
	CanReply   bool   `json:"can_reply"`    // True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	IsEnabled  bool   `json:"is_enabled"`   // True, if the connection is active
}
