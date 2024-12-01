package types

type Message struct {
	MessageID                     int32                          `json:"message_id"`                                  // Unique message identifier inside this chat
	MessageThreadID               int32                          `json:"message_thread_id,omitempty"`                 // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from,omitempty"`                              // Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                       // Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderBoostCount              int32                          `json:"sender_boost_count,omitempty"`                // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`               // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	Date                          int32                          `json:"date"`                                        // Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`            // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent from any potential bot chat which might share the same identifier.
	Chat                          *Chat                          `json:"chat"`                                        // Chat the message belongs to
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`                    // Optional. Information about the original message for forwarded messages
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`                  // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                  // Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                    // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	Quote                         *TextQuote                     `json:"quote,omitempty"`                             // Optional. For replies that quote part of the original message, the quoted part of the message
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                    // Optional. For replies to a story, the original story
	ViaBot                        *User                          `json:"via_bot,omitempty"`                           // Optional. Bot through which the message was sent
	EditDate                      int32                          `json:"edit_date,omitempty"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`             // Optional. True, if the message can't be forwarded
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`                   // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature,omitempty"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text,omitempty"`                              // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      []*MessageEntity               `json:"entities,omitempty"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`              // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	EffectID                      string                         `json:"effect_id,omitempty"`                         // Optional. Unique identifier of the message effect added to the message
	Animation                     *Animation                     `json:"animation,omitempty"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio,omitempty"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document,omitempty"`                          // Optional. Message is a general file, information about the file
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`                        // Optional. Message contains paid media; information about the paid media
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                           // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story,omitempty"`                             // Optional. Message is a forwarded story
	Video                         *Video                         `json:"video,omitempty"`                             // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice,omitempty"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption,omitempty"`                           // Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`          // Optional. True, if the caption must be shown above the message media
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`                 // Optional. True, if the message media is covered by a spoiler animation
	Contact                       *Contact                       `json:"contact,omitempty"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice,omitempty"`                              // Optional. Message is a dice with random value
	Game                          *Game                          `json:"game,omitempty"`                              // Optional. Message is a game, information about the game. More about games »
	Poll                          *Poll                          `json:"poll,omitempty"`                              // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue,omitempty"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location,omitempty"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`                    // Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`                // Optional. Service message: the group has been created
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`                  // Optional. Message is a service message about a refunded payment, information about the payment. More about payments »
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                      // Optional. Service message: users were shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                       // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`              // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                       // Optional. Service message: user boosted the chat
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`               // Optional. Service message: chat background set
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`               // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`                // Optional. Service message: forum topic edited
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`                // Optional. Service message: forum topic closed
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`              // Optional. Service message: forum topic reopened
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`        // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`      // Optional. Service message: the 'General' forum topic unhidden
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                  // Optional. Service message: a scheduled giveaway was created
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                          // Optional. The message is a scheduled giveaway message
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                  // Optional. A giveaway with public winners was completed
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`                // Optional. Service message: a giveaway without public winners was completed
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`              // Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`                // Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                  // Optional. Service message: video chat ended
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`   // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                      // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}
