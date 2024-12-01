package types

type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"` // Unique name of the time zone for which the opening hours are defined
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours"`  // List of time intervals describing business opening hours
}
