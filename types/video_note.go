package types

type VideoNote struct {
	FileID       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int32      `json:"length"`              // Video width and height (diameter of the video message) as defined by the sender
	Duration     int32      `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileSize     int64      `json:"file_size,omitempty"` // Optional. File size in bytes
}
