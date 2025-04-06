package go_nationalflooddata

type Response struct {
	Status        string         `json:"status"`
	Request       Request        `json:"request"`
	ParcelAddress *ParcelAddress `json:"parceladdress"`
	Coords        Coords         `json:"coords"`
	Result        Result         `json:"result"`
	Geocode       Geocode        `json:"geocode"`
	MatchType     *string        `json:"match_type"`
	RequestID     string         `json:"request_id"`
}
