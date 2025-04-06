package models

// Coastline represents a segment of the coastline with specific attributes.
type Coastline struct {
	// DistKm is the distance in kilometers from a reference point to the coastline.
	DistKm float64 `json:"distkm"`

	// OgcFid is the unique identifier for the feature in the Open Geospatial Consortium (OGC) format.
	OgcFid int64 `json:"ogc_fid"`
}
