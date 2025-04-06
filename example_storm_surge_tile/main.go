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

	// Retrieve storm surge tile
	category := "category1"
	z, x, y := 10, 512, 512
	stormSurgeTile, err := svc.GetStormSurgeTile(ctx, category, z, x, y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Storm Surge Tile Data: %d bytes\n", len(stormSurgeTile))
}
