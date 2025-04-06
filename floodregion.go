package go_nationalflooddata

// FloodRegion represents a flood region
type FloodRegion struct {
	// FldArID is the unique identifier for the flood area.
	FldArID string `json:"fld_ar_id"`

	// DistKm represents the distance in kilometers from a reference point to the flood area.
	DistKm float64 `json:"distkm"`

	// GeoJSON contains the geographical representation of the flood area in GeoJSON format.
	GeoJSON string `json:"geojson"`

	// ZoneSubty specifies the subtype of the flood zone, providing additional classification details.
	ZoneSubty string `json:"zone_subty"`

	// FldZone indicates the flood zone designation, which is used to assess flood risk.
	FldZone string `json:"fld_zone"`

	// DfirmID is the identifier for the Digital Flood Insurance Rate Map (DFIRM) associated with the flood area.
	DfirmID string `json:"dfirm_id"`

	// OgcFID is the unique identifier for the feature in the Open Geospatial Consortium (OGC) format.
	OgcFID int `json:"ogc_fid"`
}
