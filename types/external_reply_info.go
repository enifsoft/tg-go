package types

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`                         // Origin of the message replied to by the given message
	Chat               *Chat               `json:"chat,omitempty"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	MessageID          int32               `json:"message_id,omitempty"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Options used for link preview generation for the original message, if it is a text message
	Animation          *Animation          `json:"animation,omitempty"`            // Optional. Message is an animation, information about the animation
	Audio              *Audio              `json:"audio,omitempty"`                // Optional. Message is an audio file, information about the file
	Document           *Document           `json:"document,omitempty"`             // Optional. Message is a general file, information about the file
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Optional. Message contains paid media; information about the paid media
	Photo              []*PhotoSize        `json:"photo,omitempty"`                // Optional. Message is a photo, available sizes of the photo
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional. Message is a sticker, information about the sticker
	Story              *Story              `json:"story,omitempty"`                // Optional. Message is a forwarded story
	Video              *Video              `json:"video,omitempty"`                // Optional. Message is a video, information about the video
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional. Message is a video note, information about the video message
	Voice              *Voice              `json:"voice,omitempty"`                // Optional. Message is a voice message, information about the file
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`    // Optional. True, if the message media is covered by a spoiler animation
	Contact            *Contact            `json:"contact,omitempty"`              // Optional. Message is a shared contact, information about the contact
	Dice               *Dice               `json:"dice,omitempty"`                 // Optional. Message is a dice with random value
	Game               *Game               `json:"game,omitempty"`                 // Optional. Message is a game, information about the game. More about games »
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional. Message is a scheduled giveaway, information about the giveaway
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional. A giveaway with public winners was completed
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	Location           *Location           `json:"location,omitempty"`             // Optional. Message is a shared location, information about the location
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. Message is a native poll, information about the poll
	Venue              *Venue              `json:"venue,omitempty"`                // Optional. Message is a venue, information about the venue
}
