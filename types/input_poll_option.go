package types

type InputPollOption struct {
	Text          string           `json:"text"`                      // Option text, 1-100 characters
	TextParseMode string           `json:"text_parse_mode,omitempty"` // Optional. Mode for parsing entities in the text. See formatting options for more details. Currently, only custom emoji entities are allowed
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`   // Optional. A JSON-serialized list of special entities that appear in the poll option text. It can be specified instead of text_parse_mode
}
