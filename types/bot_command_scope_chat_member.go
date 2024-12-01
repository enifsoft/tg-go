package types

type BotCommandScopeChatMember struct {
	Type   string      `json:"type"`    // Scope type, must be chat_member
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID int64       `json:"user_id"` // Unique identifier of the target user
}
