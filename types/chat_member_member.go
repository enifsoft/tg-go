package types

type ChatMemberMember struct {
	Status    string `json:"status"`               // The member's status in the chat, always “member”
	User      *User  `json:"user"`                 // Information about the user
	UntilDate int32  `json:"until_date,omitempty"` // Optional. Date when the user's subscription will expire; Unix time
}
