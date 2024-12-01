package types

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`         // Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	WebAppName         string `json:"web_app_name,omitempty"`         // Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"` // Optional. True, if the access was granted when the bot was added to the attachment or side menu
}
