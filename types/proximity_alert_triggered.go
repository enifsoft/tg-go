package types

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"` // User that triggered the alert
	Watcher  *User `json:"watcher"`  // User that set the alert
	Distance int32 `json:"distance"` // The distance between the users
}
