package types

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`  // Type of the media, must be photo
	Media string `json:"media"` // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
}
