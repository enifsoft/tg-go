package types

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`                          // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	Creator                 *User  `json:"creator"`                              // Creator of the link
	CreatesJoinRequest      bool   `json:"creates_join_request"`                 // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool   `json:"is_primary"`                           // True, if the link is primary
	IsRevoked               bool   `json:"is_revoked"`                           // True, if the link is revoked
	Name                    string `json:"name,omitempty"`                       // Optional. Invite link name
	ExpireDate              int32  `json:"expire_date,omitempty"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             int32  `json:"member_limit,omitempty"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount int32  `json:"pending_join_request_count,omitempty"` // Optional. Number of pending join requests created using this link
	SubscriptionPeriod      int32  `json:"subscription_period,omitempty"`        // Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPrice       int32  `json:"subscription_price,omitempty"`         // Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
}
