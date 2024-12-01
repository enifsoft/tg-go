package types

type ReactionCount struct {
	Type       *ReactionType `json:"type"`        // Type of the reaction
	TotalCount int32         `json:"total_count"` // Number of times the reaction was added
}
