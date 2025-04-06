package client

// SearchType represents the type of search for the API
type SearchType string

const (
	SearchTypeAddressCoord  SearchType = "addresscoord"
	SearchTypeAddressParcel SearchType = "addressparcel"
	SearchTypeCoord         SearchType = "coord"
	SearchTypeCoordParcel   SearchType = "coordparcel"
	SearchTypePolygon       SearchType = "polygon"
)

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

// StaticMapOptions represents options for the static map query
type StaticMapOptions struct {
	Lat        float64 // Latitude for the map center
	Lng        float64 // Longitude for the map center
	Height     int     // Height of the map image
	Width      int     // Width of the map image
	ShowMarker bool    // Whether to show a marker on the map
	ShowLegend bool    // Whether to show a legend on the map
	Zoom       int     // Zoom level for the map
}
