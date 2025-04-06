package go_nationalflooddata

type Geocode struct {
	// Relevance indicates the relevance score of the geocode result, typically used to rank results.
	Relevance int `json:"relevance"`

	// MatchLevel describes the level of precision of the geocode match, such as "street" or "city".
	MatchLevel string `json:"matchLevel"`

	// Label provides a human-readable label for the geocode result, often a formatted address.
	Label string `json:"label"`

	// Latitude is the latitude coordinate of the geocode result.
	Latitude float64 `json:"latitude"`

	// Longitude is the longitude coordinate of the geocode result.
	Longitude float64 `json:"longitude"`
}
