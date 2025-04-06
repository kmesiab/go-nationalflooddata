package models

// FloodFieldHazard represents the details of a flood hazard area.
// It contains information about specific flood zones and their characteristics.
type FloodFieldHazard struct {
	// FldArID is the unique identifier for the flood area.
	FldArID string `json:"fld_ar_id"`

	// VersionID is the version identifier for the flood hazard data.
	// It indicates the version of the data being used.
	VersionID string `json:"version_id"`

	// SfhaTf indicates whether the area is a Special Flood Hazard Area (SFHA).
	// Typically, "T" for true or "F" for false.
	SfhaTf string `json:"sfha_tf"`

	// ZoneSubty is an optional field that provides the subtype of the flood zone.
	// It can be nil if no subtype is specified.
	ZoneSubty *string `json:"zone_subty"`

	// SourceCit is the source citation for the flood hazard data.
	// It provides information about the origin of the data.
	SourceCit string `json:"source_cit"`

	// FldZone is the flood zone designation, such as "AE" or "VE".
	// It indicates the level of flood risk in the area.
	FldZone string `json:"fld_zone"`

	// DfirmID is the digital FIRM ID, representing the digital version of the flood hazard map.
	DfirmID string `json:"dfirm_id"`
}
