package main

import (
	"context"
	"fmt"
	"log"

	nfd "github.com/kmesiab/go-nationalflooddata"
	"github.com/kmesiab/go-nationalflooddata/client"
)

func main() {
	svc := nfd.NewService("redacted")
	ctx := context.Background()

	// Single query
	floodData, err := svc.GetFloodData(ctx, client.FloodDataOptions{
		SearchType: client.SearchTypeAddressParcel,
		Address:    "430 Australian Ave Palm Beach, FL 33480",
		Elevation:  true,
		LOMA:       true,
		Property:   true,
		Parcel:     true,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FEMA flood zone: %+v\n", floodData.Result.FloodFldHazAr)

	// Batch
	batchResp, err := svc.GetFloodDataBatch(ctx, client.BatchDataRequest{
		Requests: []client.BatchRequest{
			{
				ID:         "req1",
				SearchType: client.SearchTypeAddressParcel,
				Address:    "430 Australian Ave Palm Beach FL 33480",
				Elevation:  true,
			},
			{
				ID:         "req2",
				SearchType: client.SearchTypeCoord,
				Lat:        "34.071783",
				Lng:        "-118.2596",
				Elevation:  false,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Batch ID: %s - poll results at: %s\n", batchResp.BatchID, batchResp.Result)
}
