package go_nationalflooddata

// StormSurge represents estimated flood water levels in the event of storms of varying categories.
type StormSurge struct {
	// Category1 is the estimated flood water level for a Category 1 storm.
	Category1 *float64 `json:"1"`

	// Category2 is the estimated flood water level for a Category 2 storm.
	Category2 *float64 `json:"2"`

	// Category3 is the estimated flood water level for a Category 3 storm.
	Category3 *float64 `json:"3"`

	// Category4 is the estimated flood water level for a Category 4 storm.
	Category4 *float64 `json:"4"`

	// Category5 is the estimated flood water level for a Category 5 storm.
	Category5 *float64 `json:"5"`
}
