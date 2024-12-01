package types

type GiveawayWinners struct {
	Chat                          *Chat   `json:"chat"`                                       // The chat that created the giveaway
	GiveawayMessageID             int32   `json:"giveaway_message_id"`                        // Identifier of the message with the giveaway in the chat
	WinnersSelectionDate          int32   `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int32   `json:"winner_count"`                               // Total number of winners in the giveaway
	Winners                       []*User `json:"winners"`                                    // List of up to 100 winners of the giveaway
	AdditionalChatCount           int32   `json:"additional_chat_count,omitempty"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PrizeStarCount                int32   `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int32   `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	UnclaimedPrizeCount           int32   `json:"unclaimed_prize_count,omitempty"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                bool    `json:"only_new_members,omitempty"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   bool    `json:"was_refunded,omitempty"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              string  `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
}
