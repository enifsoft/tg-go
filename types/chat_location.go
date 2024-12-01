package types

type ChatLocation struct {
	Location *Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string    `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}
