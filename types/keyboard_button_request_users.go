package types

type KeyboardButtonRequestUsers struct {
	RequestID       int32 `json:"request_id"`                 // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	UserIsBOT       bool  `json:"user_is_bot,omitempty"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   bool  `json:"user_is_premium,omitempty"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     int32 `json:"max_quantity,omitempty"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool  `json:"request_name,omitempty"`     // Optional. Pass True to request the users' first and last names
	RequestUsername bool  `json:"request_username,omitempty"` // Optional. Pass True to request the users' usernames
	RequestPhoto    bool  `json:"request_photo,omitempty"`    // Optional. Pass True to request the users' photos
}
