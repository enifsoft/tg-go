package types

type Animation struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int32      `json:"width"`               // Video width as defined by the sender
	Height       int32      `json:"height"`              // Video height as defined by the sender
	Duration     int32      `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Animation thumbnail as defined by the sender
	FileName     string     `json:"file_name,omitempty"` // Optional. Original animation filename as defined by the sender
	MimeType     string     `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
}
