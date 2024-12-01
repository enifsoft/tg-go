package types

type MessageOriginChat struct {
	Type            string `json:"type"`                       // Type of the message origin, always “chat”
	Date            int32  `json:"date"`                       // Date the message was sent originally in Unix time
	SenderChat      *Chat  `json:"sender_chat"`                // Chat that sent the message originally
	AuthorSignature string `json:"author_signature,omitempty"` // Optional. For messages originally sent by an anonymous chat administrator, original message author signature
}
