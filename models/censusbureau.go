package models

// CensusBureau represents data related to the U.S. Census Bureau, including census blocks and statistical areas.
type CensusBureau struct {
	// CensusBlock is the identifier for the census block, which is a geographic area used by the U.S. Census Bureau.
	CensusBlock string `json:"census_block"`

	// CBSA is a pointer to a CBSA struct, representing a Core-Based Statistical Area associated with the census block.
	CBSA *CBSA `json:"cbsa"`

	// MetDiv is a pointer to a MetDiv struct, representing a Metropolitan Division associated with the census block.
	MetDiv *MetDiv `json:"metdiv"`
}
