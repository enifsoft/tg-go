package types

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`       // Type of the background, always “chat_theme”
	ThemeName string `json:"theme_name"` // Name of the chat theme, which is usually an emoji
}
