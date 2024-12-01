package types

type ChatMember struct {
	Status      string `json:"status"`                 // The member's status in the chat, always “creator”
	User        *User  `json:"user"`                   // Information about the user
	IsAnonymous bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CustomTitle string `json:"custom_title,omitempty"` // Optional. Custom title for this user
}
