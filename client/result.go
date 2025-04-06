package client

import "github.com/kmesiab/go-nationalflooddata/models"

// Result contains FEMA flood data for a location.
type Result struct {
	FloodFirmPan  []models.FloodFirmPan     `json:"flood.s_firm_pan"`
	FloodFldHazAr []models.FloodFieldHazard `json:"flood.s_fld_haz_ar"`
	FloodPolAr    []models.FloodPolAr       `json:"flood.s_pol_ar"`
	CensusBureau  *models.CensusBureau      `json:"census_bureau,omitempty"`
	Community     *models.Community         `json:"community,omitempty"`
	Elevation     *models.Elevation         `json:"elevation,omitempty"`
	Property      *models.Property          `json:"property,omitempty"`
	Loma          *[]models.Loma            `json:"loma,omitempty"`
	Geocode       *models.Geocode           `json:"geocode,omitempty"`
	DeniedAccess  []string
}
