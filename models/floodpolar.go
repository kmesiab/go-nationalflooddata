package models

// FloodPolAr represents a flood policy area with specific identifiers and names.
type FloodPolAr struct {
	// CommNo is the community number associated with the flood policy area.
	CommNo string `json:"comm_no"`

	// PolName1 is the primary name of the flood policy area.
	PolName1 string `json:"pol_name1"`

	// CoFips is the FIPS (Federal Information Processing Standards) code for the county.
	CoFips string `json:"co_fips"`

	// CID is the community identifier for the flood policy area.
	CID string `json:"cid"`

	// ComNfoID is the community information ID, which provides additional details about the community.
	ComNfoID string `json:"com_nfo_id"`

	// PolArID is the unique identifier for the flood policy area.
	PolArID string `json:"pol_ar_id"`
}
