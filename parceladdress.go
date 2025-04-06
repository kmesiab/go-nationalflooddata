package go_nationalflooddata

// ParcelAddress represents address data for a parcel
type ParcelAddress struct {
	// AddrNumber is the street number of the parcel's address.
	AddrNumber string `json:"addr_number"`

	// AddrStreetName is the name of the street for the parcel's address.
	AddrStreetName string `json:"addr_street_name"`

	// AddrStreetPrefix is an optional prefix for the street name, such as "N" for North.
	AddrStreetPrefix *string `json:"addr_street_prefix"`

	// AddrStreetSuffix is an optional suffix for the street name, such as "NW" for Northwest.
	AddrStreetSuffix *string `json:"addr_street_suffix"`

	// AddrStreetType is the type of the street, such as "Ave" for Avenue or "St" for Street.
	AddrStreetType string `json:"addr_street_type"`

	// CountyId is the identifier for the county where the parcel is located.
	CountyId string `json:"county_id"`

	// CountyName is the name of the county where the parcel is located.
	CountyName string `json:"county_name"`

	// MuniName is the name of the municipality where the parcel is located.
	MuniName string `json:"muni_name"`

	// ParcelId is the unique identifier for the parcel.
	ParcelId string `json:"parcel_id"`

	// Physcity is the physical city where the parcel is located.
	Physcity string `json:"physcity"`

	// Physzip is the physical ZIP code for the parcel's location.
	Physzip string `json:"physzip"`

	// StateAbbr is the abbreviation of the state where the parcel is located.
	StateAbbr string `json:"state_abbr"`
}
