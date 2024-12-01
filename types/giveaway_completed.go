package types

type GiveawayCompleted struct {
	WinnerCount         int32    `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount int32    `json:"unclaimed_prize_count,omitempty"` // Optional. Number of undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`      // Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
}
