package types

type InputMedia struct {
	Type                  string           `json:"type"`                               // Type of the result, must be photo
	Media                 string           `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the photo needs to be covered with a spoiler animation
}
