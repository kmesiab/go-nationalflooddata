package client

import "github.com/kmesiab/go-nationalflooddata/models"

// FloodMapContent represents raw flood map data
type FloodMapContent struct {
	Result FloodMapContentResult `json:"result"`
}

// FloodMapContentResult contains raw flood map data
type FloodMapContentResult struct {
	// BFEList is a list of Base Flood Elevation (BFE) items,
	// which provide information about the elevation of floodwaters
	// during a base flood event.
	BFEList []models.BFEListItem `json:"bfelist"`

	// FloodRegions is a list of FloodRegion items, each representing a specific
	// flood region with associated data such as flood zone, distance, and
	// geographical information.
	FloodRegions []models.FloodRegion `json:"floodregions"`
}
