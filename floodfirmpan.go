package go_nationalflooddata

// FloodFirmPan represents the details of a flood insurance rate map (FIRM) panel.
// It contains information about the specific panel used in flood mapping.
type FloodFirmPan struct {
	// Suffix is the suffix of the FIRM panel identifier.
	Suffix string `json:"suffix"`

	// PnpReason is an optional field that provides the reason for the
	// panel not printed (PNP) status. It can be nil if the panel is printed.
	PnpReason *string `json:"pnp_reason"`

	// FirmPan is the unique identifier for the FIRM panel.
	FirmPan string `json:"firm_pan"`

	// EffDate is the effective date of the FIRM panel.
	// It indicates when the panel became effective for flood insurance purposes.
	EffDate string `json:"eff_date"`

	// FirmID is the identifier for the FIRM panel within the flood insurance study.
	FirmID string `json:"firm_id"`

	// DfirmID is the digital FIRM ID, representing the digital version of the FIRM panel.
	DfirmID string `json:"dfirm_id"`

	// StFips is the state FIPS (Federal Information Processing Standards) code.
	// It identifies the state associated with the FIRM panel.
	StFips string `json:"st_fips"`

	// PanelTyp describes the type of panel, such as "Countywide, Panel Printed".
	// It provides additional context about the panel's scope and format.
	PanelTyp string `json:"panel_typ"`

	// Panel is the panel number within the FIRM.
	// It is used to identify the specific section of the map.
	Panel string `json:"panel"`
}
