package go_nationalflooddata

// Result contains FEMA flood data for a location.
type Result struct {
	FloodFirmPan  []FloodFirmPan     `json:"flood.s_firm_pan"`
	FloodFldHazAr []FloodFieldHazard `json:"flood.s_fld_haz_ar"`
	FloodPolAr    []FloodPolAr       `json:"flood.s_pol_ar"`
	CensusBureau  *CensusBureau      `json:"census_bureau,omitempty"`
	Community     *Community         `json:"community,omitempty"`
	Elevation     *Elevation         `json:"elevation,omitempty"`
	Property      *Property          `json:"property,omitempty"`
	Loma          *[]Loma            `json:"loma,omitempty"`
	Geocode       *Geocode           `json:"geocode,omitempty"`
	DeniedAccess  []string
}
