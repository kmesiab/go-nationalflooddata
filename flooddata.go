package go_nationalflooddata

// FloodData represents the response from the FEMA Flood Data API, containing detailed information about a specific location.
type FloodData struct {
	// ParcelAddress contains the address details of the parcel.
	ParcelAddress ParcelAddress `json:"parceladdress"`

	// Geocode provides the geocoding information for the location.
	Geocode Geocode `json:"geocode"`

	// Coords holds the latitude and longitude coordinates of the location.
	Coords Coords `json:"coords"`

	// Result contains the flood-related data and analysis results for the location.
	Result Result `json:"result"`

	// MatchType indicates the type of match found for the location query, such as address or coordinates.
	MatchType *string `json:"match_type"`
}

// FloodDataBatch represents a batch response from the FEMA Flood Data API, used for processing multiple requests at once.
type FloodDataBatch struct {
	// BatchID is the unique identifier for the batch request.
	BatchID string `json:"batch_id"`

	// Result is a presigned URL for an S3 object containing the batch result data.
	Result string `json:"result"`
}
