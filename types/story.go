package types

type Story struct {
	Chat *Chat `json:"chat"` // Chat that posted the story
	ID   int32 `json:"id"`   // Unique identifier for the story in the chat
}
