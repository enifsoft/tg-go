package types

type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`                                  // Chat the user belongs to
	From                    *User           `json:"from"`                                  // Performer of the action, which resulted in the change
	Date                    int32           `json:"date"`                                  // Date the change was done in Unix time
	OldChatMember           *ChatMember     `json:"old_chat_member"`                       // Previous information about the chat member
	NewChatMember           *ChatMember     `json:"new_chat_member"`                       // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
}
