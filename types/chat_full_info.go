package types

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`                                                // Unique identifier for this chat
	Type                               string                `json:"type"`                                              // Type of the chat: "private", "group", "supergroup", or "channel"
	AccentColorID                      int32                 `json:"accent_color_id"`                                   // Identifier of the accent color
	MaxReactionCount                   int32                 `json:"max_reaction_count"`                                // Maximum number of reactions per message
	Title                              string                `json:"title,omitempty"`                                   // Optional. Title for supergroups, channels, and group chats
	Username                           string                `json:"username,omitempty"`                                // Optional. Username for private chats, supergroups, and channels
	FirstName                          string                `json:"first_name,omitempty"`                              // Optional. First name in a private chat
	LastName                           string                `json:"last_name,omitempty"`                               // Optional. Last name in a private chat
	IsForum                            bool                  `json:"is_forum,omitempty"`                                // Optional. True if the supergroup chat is a forum
	Photo                              *ChatPhoto            `json:"photo,omitempty"`                                   // Optional. Chat photo
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`                        // Optional. List of active chat usernames
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`                               // Optional. Birthdate in private chats
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`                          // Optional. Business intro in private chats
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`                       // Optional. Business location in private chats
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`                  // Optional. Business opening hours in private chats
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`                           // Optional. Personal channel in private chats
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`                     // Optional. Allowed reactions in the chat
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`              // Optional. Custom emoji ID for background
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`                 // Optional. Accent color ID for profile background
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`      // Optional. Custom emoji ID for profile background
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`            // Optional. Custom emoji ID for emoji status
	EmojiStatusExpirationDate          int                   `json:"emoji_status_expiration_date,omitempty"`            // Optional. Expiration date for emoji status (Unix time)
	Bio                                string                `json:"bio,omitempty"`                                     // Optional. Bio in private chats
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`                    // Optional. True if privacy settings restrict forwards
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"` // Optional. True if voice/video messages are restricted
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`                   // Optional. True if users must join to send messages
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`                         // Optional. True if users need approval to join
	Description                        string                `json:"description,omitempty"`                             // Optional. Chat description
	InviteLink                         string                `json:"invite_link,omitempty"`                             // Optional. Primary invite link
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`                          // Optional. Most recent pinned message
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`                             // Optional. Default chat member permissions
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`                     // Optional. True if paid media messages can be sent
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`                         // Optional. Minimum delay between messages for unprivileged users (in seconds)
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`                  // Optional. Boost count to ignore slow mode
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`                // Optional. Auto-delete time for messages (in seconds)
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`        // Optional. True if aggressive anti-spam is enabled
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`                      // Optional. True if only admins and bots are visible
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`                   // Optional. True if messages can't be forwarded
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`                     // Optional. True if new members can see old messages
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`                        // Optional. Group sticker set name
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`                     // Optional. True if the bot can change the sticker set
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`           // Optional. Custom emoji sticker set name
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`                          // Optional. Linked chat ID
	Location                           *ChatLocation         `json:"location,omitempty"`                                // Optional. Chat location
}
