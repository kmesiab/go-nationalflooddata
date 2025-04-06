package go_nationalflooddata

type MetDiv struct {
	// Metdivfp is likely the Federal Information Processing Standards (FIPS) code for the metropolitan division.
	// FIPS codes are used to uniquely identify geographic areas.
	Metdivfp string `json:"metdivfp"`

	// Name is the name of the metropolitan division, providing a human-readable identifier for the area.
	Name string `json:"name"`
}
