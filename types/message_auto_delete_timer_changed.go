package types

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int32 `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}
