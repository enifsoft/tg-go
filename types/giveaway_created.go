package types

type GiveawayCreated struct {
	PrizeStarCount int32 `json:"prize_star_count,omitempty"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
}
