package go_nationalflooddata

// BaseFloodElevation represents the details of a base flood elevation (BFE) area.
type BaseFloodElevation struct {
	// BfeLnID is the unique identifier for the base flood elevation line.
	BfeLnID *string `json:"bfe_ln_id"`

	// BfeType indicates the type of base flood elevation.
	BfeType string `json:"bfe_type"`

	// DfirmID is the identifier for the Digital Flood Insurance Rate Map (DFIRM).
	DfirmID string `json:"dfirm_id"`

	// DistKm is the distance in kilometers from a reference point to the BFE line.
	DistKm int `json:"distkm"`

	// Elevation is the elevation value of the base flood elevation.
	Elevation string `json:"elevation"`

	// FldArID is the identifier for the flood area.
	FldArID string `json:"fld_ar_id"`

	// FldZone is the flood zone designation.
	FldZone string `json:"fld_zone"`

	// LenUnit is the unit of measurement for length, such as feet or meters.
	LenUnit string `json:"len_unit"`

	// VDatum is the vertical datum used for the elevation measurement.
	VDatum string `json:"v_datum"`

	// ZoneSubty is an optional subtype for the flood zone.
	ZoneSubty *string `json:"zone_subty"`
}

// BFEListItem represents a base flood elevation line item with additional details.
type BFEListItem struct {
	// BfeLnID is the unique identifier for the base flood elevation line.
	BfeLnID string `json:"bfe_ln_id"`

	// VDatum is the vertical datum used for the elevation measurement.
	VDatum string `json:"v_datum"`

	// DistKm is the distance in kilometers from a reference point to the BFE line.
	DistKm float64 `json:"distkm"`

	// VersionID is the version identifier for the BFE data.
	VersionID string `json:"version_id"`

	// SourceCit is the source citation for the BFE data.
	SourceCit string `json:"source_cit"`

	// GeoJSON is the GeoJSON representation of the BFE line.
	GeoJSON string `json:"geojson"`

	// Elev is the elevation value of the base flood elevation.
	Elev int `json:"elev"`

	// DfirmID is the identifier for the Digital Flood Insurance Rate Map (DFIRM).
	DfirmID string `json:"dfirm_id"`

	// LenUnit is the unit of measurement for length, such as feet or meters.
	LenUnit string `json:"len_unit"`

	// OgcFID is the unique identifier for the feature in the Open Geospatial Consortium (OGC) format.
	OgcFID int `json:"ogc_fid"`
}
