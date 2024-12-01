package types

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`               // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`    // Optional. True, if private chats with users can be chosen
	AllowBOTChats     bool   `json:"allow_bot_chats,omitempty"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"` // Optional. True, if channel chats can be chosen
}
