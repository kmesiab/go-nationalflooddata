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

	// Query for raw flood map data
	floodMapContent, err := svc.GetFloodMapRaw(ctx, client.FloodMapRawOptions{
		Lat:       34.071783,
		Lng:       -118.2596,
		Size:      0.08,
		GeoJSON:   true,
		ExcludeX:  true,
		Elevation: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Flood Map Content: %+v\n", floodMapContent)

	// Query for static flood map image
	staticMap, err := svc.GetStaticFloodMap(ctx, client.StaticMapOptions{
		Lat:        34.071783,
		Lng:        -118.2596,
		Height:     600,
		Width:      800,
		ShowMarker: true,
		ShowLegend: true,
		Zoom:       13,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Static Map Image Data: %d bytes\n", len(staticMap))
}
