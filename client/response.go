package client

import "github.com/kmesiab/go-nationalflooddata/models"

type Response struct {
	Status        string                `json:"status"`
	Request       Request               `json:"request"`
	ParcelAddress *models.ParcelAddress `json:"parceladdress"`
	Coords        models.Coords         `json:"coords"`
	Result        Result                `json:"result"`
	Geocode       models.Geocode        `json:"geocode"`
	MatchType     *string               `json:"match_type"`
	RequestID     string                `json:"request_id"`
}
