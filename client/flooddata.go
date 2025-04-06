package client

import "github.com/kmesiab/go-nationalflooddata/models"

// FloodData represents the response from the FEMA Flood Data API, containing detailed information about a specific location.
type FloodData struct {
	// ParcelAddress contains the address details of the parcel.
	ParcelAddress models.ParcelAddress `json:"parceladdress"`

	// Geocode provides the geocoding information for the location.
	Geocode models.Geocode `json:"geocode"`

	// Coords holds the latitude and longitude coordinates of the location.
	Coords models.Coords `json:"coords"`

	// Result contains the flood-related data and analysis results for the location.
	Result Result `json:"result"`

	// MatchType indicates the type of match found for the location query, such as address or coordinates.
	MatchType *string `json:"match_type"`
}
