package types

type BotCommandScopeChatAdministrators struct {
	Type   string      `json:"type"`    // Scope type, must be chat_administrators
	ChatID interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}
