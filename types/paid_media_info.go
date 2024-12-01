package types

type PaidMediaInfo struct {
	StarCount int32        `json:"star_count"` // The number of Telegram Stars that must be paid to buy access to the media
	PaidMedia []*PaidMedia `json:"paid_media"` // Information about the paid media
}
