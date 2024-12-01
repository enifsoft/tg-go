package types

type RefundedPayment struct {
	Currency                string `json:"currency"`                             // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars. Currently, always “XTR”
	TotalAmount             int32  `json:"total_amount"`                         // Total refunded price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string `json:"invoice_payload"`                      // Bot-specified invoice payload
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`           // Telegram payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"` // Optional. Provider payment identifier
}
