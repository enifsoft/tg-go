package types

type PassportFile struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int64  `json:"file_size"`      // File size in bytes
	FileDate     int32  `json:"file_date"`      // Unix time when the file was uploaded
}
