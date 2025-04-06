package models

type Waterbody struct {
	// AreaSqKm is the area of the waterbody in square kilometers.
	AreaSqKm string `json:"areasqkm"`

	// DistKm is the distance in kilometers from a reference point to the waterbody.
	DistKm float64 `json:"distkm"`

	// GnisID is the Geographic Names Information System (GNIS) identifier for the waterbody.
	GnisID string `json:"gnis_id"`

	// Name is the name of the waterbody.
	Name string `json:"name"`

	// ObjectID is a unique identifier for the waterbody object.
	ObjectID string `json:"objectid"`

	// OgcFid is the unique identifier for the feature in the Open Geospatial Consortium (OGC) format.
	OgcFid int64 `json:"ogc_fid"`

	// State is the state where the waterbody is located.
	State string `json:"state"`
}
