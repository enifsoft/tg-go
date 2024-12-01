package types

type VideoChatScheduled struct {
	StartDate int32 `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}
