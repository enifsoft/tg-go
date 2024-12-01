package types

type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always “restricted”
	User                  *User  `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	CanSendAudios         bool   `json:"can_send_audios"`           // True, if the user is allowed to send audios
	CanSendDocuments      bool   `json:"can_send_documents"`        // True, if the user is allowed to send documents
	CanSendPhotos         bool   `json:"can_send_photos"`           // True, if the user is allowed to send photos
	CanSendVideos         bool   `json:"can_send_videos"`           // True, if the user is allowed to send videos
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`      // True, if the user is allowed to send video notes
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`      // True, if the user is allowed to send voice notes
	CanSendPolls          bool   `json:"can_send_polls"`            // True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True, if the user is allowed to add web page previews to their messages
	CanChangeInfo         bool   `json:"can_change_info"`           // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users"`          // True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages"`          // True, if the user is allowed to pin messages
	CanManageTopics       bool   `json:"can_manage_topics"`         // True, if the user is allowed to create forum topics
	UntilDate             int32  `json:"until_date"`                // Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever
}
