package types

type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int32  `json:"value"` // Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
}
