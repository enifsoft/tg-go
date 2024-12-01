package types

type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always “left”
	User   *User  `json:"user"`   // Information about the user
}
