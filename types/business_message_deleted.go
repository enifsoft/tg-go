package types

type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 *Chat   `json:"chat"`                   // Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	MessageIds           []int32 `json:"message_ids"`            // The list of identifiers of deleted messages in the chat of the business account
}
