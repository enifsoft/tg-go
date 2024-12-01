package types

type UsersShared struct {
	RequestID int32         `json:"request_id"` // Identifier of the request
	Users     []*SharedUser `json:"users"`      // Information about users shared with the bot.
}
