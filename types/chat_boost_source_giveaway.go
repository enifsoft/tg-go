package types

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`                     // Source of the boost, always “giveaway”
	GiveawayMessageID int32  `json:"giveaway_message_id"`        // Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	User              *User  `json:"user,omitempty"`             // Optional. User that won the prize in the giveaway if any; for Telegram Premium giveaways only
	PrizeStarCount    int32  `json:"prize_star_count,omitempty"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`     // Optional. True, if the giveaway was completed, but there was no user to win the prize
}
