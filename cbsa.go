package go_nationalflooddata

// CBSA represents a Core-Based Statistical Area, which is a U.S. geographic area defined by the Office of Management and Budget.
type CBSA struct {
	// Cbsafp is the unique identifier for the Core-Based Statistical Area.
	Cbsafp string `json:"cbsafp"`

	// Name is the name of the Core-Based Statistical Area.
	Name string `json:"name"`
}
