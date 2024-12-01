package types

type PollOption struct {
	Text         string           `json:"text"`                    // Option text, 1-100 characters
	TextEntities []*MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount   int32            `json:"voter_count"`             // Number of users that voted for this option
}
