package models

// Elevation represents various elevation-related data for a property, including flood and storm surge information.
type Elevation struct {
	// PropertyElevation is the elevation of the property in meters.
	PropertyElevation float64 `json:"propertyelevation"`

	// FloodBaseFloodElevation is a list of base flood elevation data associated with the property.
	FloodBaseFloodElevation []BaseFloodElevation `json:"flood.basefloodelevation"`

	// Coastline is a list of coastline segments related to the property's location.
	Coastline []Coastline `json:"coastline"`

	// Waterbody is a list of waterbodies near the property.
	Waterbody []Waterbody `json:"waterbody"`

	// StormSurge contains estimated flood water levels for different storm categories.
	StormSurge StormSurge `json:"stormsurge"`
}
