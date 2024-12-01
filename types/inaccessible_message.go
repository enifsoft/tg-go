package types

type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`       // Chat the message belonged to
	MessageID int32 `json:"message_id"` // Unique message identifier inside the chat
	Date      int32 `json:"date"`       // Always 0. The field can be used to differentiate regular and inaccessible messages.
}
