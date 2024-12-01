package types

type BusinessLocation struct {
	Address  string    `json:"address"`            // Address of the business
	Location *Location `json:"location,omitempty"` // Optional. Location of the business
}
