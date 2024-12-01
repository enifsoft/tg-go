package types

type Giveaway struct {
	Chats                         []*Chat  `json:"chats"`                                      // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int32    `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int32    `json:"winner_count"`                               // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`                 // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              string   `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PrizeStarCount                int32    `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount int32    `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
}
