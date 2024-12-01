package types

type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"` // New members that were invited to the video chat
}
