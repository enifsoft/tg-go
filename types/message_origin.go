package types

type MessageOrigin struct {
	Type       string `json:"type"`        // Type of the message origin, always “user”
	Date       int32  `json:"date"`        // Date the message was sent originally in Unix time
	SenderUser *User  `json:"sender_user"` // User that sent the message originally
}
