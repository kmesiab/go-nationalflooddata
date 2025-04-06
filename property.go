package go_nationalflooddata

// Property represents property data, including various attributes related to the structure and use of the property.
type Property struct {
	// SqFt is the square footage of the property, represented as a string.
	SqFt *string `json:"sqft"`

	// YearBuilt is the year the property was built, represented as a string.
	YearBuilt *string `json:"yearbuilt"`

	// PropertyUseDescription provides a description of how the property is used, such as residential or commercial.
	PropertyUseDescription *string `json:"propertyusedescription"`

	// ConstructionDesc describes the construction type or materials used for the property.
	ConstructionDesc *string `json:"constructiondesc"`

	// StoriesCount is the number of stories or levels in the property, represented as a string.
	StoriesCount *string `json:"storiescount"`

	// FireResistance indicates the fire resistance rating or characteristics of the property.
	FireResistance *string `json:"fireresistance"`

	// ParkingGarageType describes the type of parking garage associated with the property, if any.
	ParkingGarageType *string `json:"parkinggaragetype"`

	// ParkingGarageArea is the area of the parking garage, represented as a string.
	ParkingGarageArea *string `json:"parkinggaragearea"`
}
