package models

// Community represents information about a community's participation in the National Flood Insurance Program (NFIP).
type Community struct {
	// Firm is the Flood Insurance Rate Map (FIRM) identifier for the community.
	Firm string `json:"firm"`

	// RegemerSanction is the date of the community's regular emergency sanction, returned as unformatted text.
	RegemerSanction string `json:"regemer_sanction"`

	// Tribal indicates whether the community is identified as Tribal in the NFIP Community Status Book.
	Tribal string `json:"tribal"`

	// Notes are optional notes related to the community's flood insurance status.
	Notes *string `json:"notes"`

	// CommName is the name of the community.
	CommName string `json:"comm_name"`

	// CommPart indicates whether the community participates in the NFIP.
	CommPart bool `json:"comm_part"`

	// Fhbm is the date of the first flood hazard boundary map for the community, returned as unformatted text.
	Fhbm string `json:"fhbm"`

	// Curreff is the date the current FIRM became effective, returned as unformatted text. It may include additional information such as (>) for future dates, (M) for "No elevation determined", (S) for "Suspended Community", or (E) for "Indicates Entry in Emergency Program".
	Curreff string `json:"curreff"`
}
