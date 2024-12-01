package types

type MessageOriginChannel struct {
	Type            string `json:"type"`                       // Type of the message origin, always “channel”
	Date            int32  `json:"date"`                       // Date the message was sent originally in Unix time
	Chat            *Chat  `json:"chat"`                       // Channel chat to which the message was originally sent
	MessageID       int32  `json:"message_id"`                 // Unique message identifier inside the chat
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. Signature of the original post author
}
