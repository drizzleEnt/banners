package main

import (
	"context"
	"log"

	"github.com/drizzleent/banners/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init app %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %v", err)
	}

}
