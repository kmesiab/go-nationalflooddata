package models

// Loma represents a Letter of Map Amendment (LOMA) record.
// It contains information about amendments to flood maps for specific properties.
type Loma struct {
	// CaseNumber is the unique identifier for the LOMA case.
	CaseNumber string `json:"casenumber"`

	// CID is the community identifier associated with the LOMA.
	CID string `json:"cid"`

	// CommunityN is the name of the community where the LOMA is applicable.
	CommunityN string `json:"communityn"`

	// DateEnded is the date when the LOMA case was concluded.
	DateEnded string `json:"dateended"`

	// Determinat is the determination letter type for the LOMA.
	Determinat string `json:"determinat"`

	// Lat is the latitude coordinate of the property related to the LOMA.
	Lat string `json:"lat"`

	// Lon is the longitude coordinate of the property related to the LOMA.
	Lon string `json:"lon"`

	// Miles is the distance in miles from a reference point to the property.
	Miles float64 `json:"miles"`

	// PdfHyperli is the hyperlink identifier for the LOMA PDF document.
	PdfHyperli string `json:"pdfhyperli"`

	// PdfLink is the URL link to download the LOMA PDF document.
	PdfLink string `json:"pdflink"`

	// ProjectCat is the category of the project associated with the LOMA.
	ProjectCat string `json:"projectcat"`

	// ProjectName is the name of the project associated with the LOMA.
	ProjectName string `json:"projectnam"`

	// Status indicates the current status of the LOMA case.
	Status string `json:"status"`
}
