package go_nationalflooddata

// Request represents a flood data query request, containing various parameters for searching flood data.
type Request struct {
	// Searchtype specifies the type of search to be performed, such as by address or coordinates.
	Searchtype string `json:"searchtype"`

	// Address is the address to be used for the flood data query.
	Address string `json:"address"`

	// Lat is the latitude coordinate for the flood data query.
	Lat string `json:"lat"`

	// Lng is the longitude coordinate for the flood data query.
	Lng string `json:"lng"`

	// Apn is the Assessor's Parcel Number, which can be used for querying flood data.
	Apn interface{} `json:"apn"`

	// County is the county information, which can be used for querying flood data.
	County interface{} `json:"county"`

	// State is the state information, which can be used for querying flood data.
	State interface{} `json:"state"`

	// MatchType indicates the type of match found for the query, such as exact or partial.
	MatchType string `json:"match_type"`
}

// BatchRequest represents a batch request item, allowing multiple flood data queries to be processed together.
type BatchRequest struct {
	// ID is the unique identifier for the batch request item.
	ID string `json:"id"`

	// SearchType specifies the type of search to be performed for the batch request.
	SearchType SearchType `json:"searchtype"`

	// Address is the address to be used for the batch request query, if applicable.
	Address string `json:"address,omitempty"`

	// Lat is the latitude coordinate for the batch request query, if applicable.
	Lat string `json:"lat,omitempty"`

	// Lng is the longitude coordinate for the batch request query, if applicable.
	Lng string `json:"lng,omitempty"`

	// Polygon is the polygon data for the batch request query, if applicable.
	Polygon string `json:"polygon,omitempty"`

	// LOMA indicates whether a Letter of Map Amendment (LOMA) is requested.
	LOMA bool `json:"loma,omitempty"`

	// Elevation indicates whether elevation data is requested.
	Elevation bool `json:"elevation,omitempty"`

	// Property indicates whether property data is requested.
	Property bool `json:"property,omitempty"`

	// Parcel indicates whether parcel data is requested.
	Parcel bool `json:"parcel,omitempty"`
}

// BatchDataRequest represents the full batch request, containing multiple batch request items.
type BatchDataRequest struct {
	// APIKey is the API key used for authentication with the flood data service.
	APIKey string `json:"apiKey"`

	// Requests is a list of batch request items to be processed together.
	Requests []BatchRequest `json:"requests"`
}
