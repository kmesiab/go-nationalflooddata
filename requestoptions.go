package go_nationalflooddata

// SearchType represents the type of search for the API
type SearchType string

// FloodDataOptions represents options for the flood data query
type FloodDataOptions struct {
	SearchType SearchType
	Address    string
	Lat        float64
	Lng        float64
	Polygon    string
	LOMA       bool
	Elevation  bool
	Property   bool
	Parcel     bool
}

// FloodMapRawOptions represents options for the flood map raw query
type FloodMapRawOptions struct {
	Lat       float64
	Lng       float64
	Size      float64
	GeoJSON   bool
	ExcludeX  bool
	Elevation bool
}
