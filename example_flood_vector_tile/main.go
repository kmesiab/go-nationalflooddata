package main

import (
	"context"
	"fmt"
	"log"

	nfd "github.com/kmesiab/go-nationalflooddata"
)

func main() {
	svc := nfd.NewService("redacted")
	ctx := context.Background()

	// Retrieve flood vector tile
	z, x, y := 10, 512, 512
	vectorTile, err := svc.GetFloodVectorTile(ctx, z, x, y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Flood Vector Tile Data: %d bytes\n", len(vectorTile))
}
