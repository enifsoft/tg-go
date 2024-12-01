package types

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always “hidden_user”
	Date           int32  `json:"date"`             // Date the message was sent originally in Unix time
	SenderUserName string `json:"sender_user_name"` // Name of the user that sent the message originally
}
