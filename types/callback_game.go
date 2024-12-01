package types

type CallbackGame struct {
	UserID             int64  `json:"user_id"`                        // Yes
	Score              int32  `json:"score"`                          // Yes
	Force              bool   `json:"force,omitempty"`                // Optional
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Optional
	ChatID             int64  `json:"chat_id,omitempty"`              // Optional
	MessageID          int32  `json:"message_id,omitempty"`           // Optional
	InlineMessageID    string `json:"inline_message_id,omitempty"`    // Optional
}
