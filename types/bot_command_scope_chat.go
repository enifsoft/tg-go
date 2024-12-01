package types

type BotCommandScopeChat struct {
	Type   string      `json:"type"`    // Scope type, must be chat
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}
