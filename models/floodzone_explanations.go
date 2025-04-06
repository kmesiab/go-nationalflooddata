package models

// FloodZoneExplanation represents the explanation for different flood zone designations.
type FloodZoneExplanation map[string]string

// FloodZoneExplanations is a lookup table for flood zone codes and their explanations.
var FloodZoneExplanations = FloodZoneExplanation{
	"A":          "An area with a 1% annual chance of flood; does not have base flood elevations (BFEs) available.",
	"AE":         "An area with a 1% annual chance of flood; base flood elevations BFEs are available.",
	"AH":         "An area with a 1% annual chance of flood with flood depths ranging from 1 to 3 feet, generally near pond or pooling areas. BFEs are available.",
	"AO":         "An area with a 1% annual chance of flood with flood depths ranging from 1 to 3 feet, generally sheet flow on sloping terrain. BFEs are available.",
	"AR":         "An area inundated by flooding, for which BFEs or average depths have been determined. This is an area that was previously, and will again, be protected from the 1% annual chance flood by a Federal flood protection system whose restoration is Federally funded and underway.",
	"A1-A30":     "An area with a 1% annual chance flooding, for which BFEs have been determined.",
	"B, X500":    "An area with at least a 0.2% chance of annual flood or with a 1% annual chance of flood with average depths less than one foot or with drainage area less than one square mile. (C is the older designation and X500 is the current designation.)",
	"C, X":       "An area outside the 0.2% and 1% annual chance of flood regions. (C is the older designation and X500 is the current designation.)",
	"D":          "An area where flooding is possible but has not been studied.",
	"V":          "An area with a 1% annual chance flooding with velocity hazard due to waves; BFEs have are not available.",
	"VE, V1-V30": "An area with a 1% annual chance flooding with velocity hazard due to waves; BFEs have are available.",
}
