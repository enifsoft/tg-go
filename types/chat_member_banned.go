package types

type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always “kicked”
	User      *User  `json:"user"`       // Information about the user
	UntilDate int32  `json:"until_date"` // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
}
