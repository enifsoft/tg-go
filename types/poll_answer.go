package types

type PollAnswer struct {
	PollID    string  `json:"poll_id"`              // Unique poll identifier
	VoterChat *Chat   `json:"voter_chat,omitempty"` // Optional. The chat that changed the answer to the poll, if the voter is anonymous
	User      *User   `json:"user,omitempty"`       // Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	OptionIds []int32 `json:"option_ids"`           // 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}
