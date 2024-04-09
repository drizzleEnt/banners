package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/drizzleent/banners/internal/app"
	"github.com/drizzleent/banners/pkg/client/db"
)

func main() {
	ctx := context.Background()
	app, _ := app.New(ctx)
	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)
		ctx := context.Background()
		url := gofakeit.URL()
		text := gofakeit.Sentence(rand.Intn((7 - 3) + 3))
		title := gofakeit.BookTitle()
		feature := rand.Intn((20 - 1) + 1)
		tagCount := rand.Intn((5 - 1) + 1)
		tags := make([]int, tagCount)
		for i := 0; i < tagCount; i++ {
			tags[i] = rand.Intn((10 - 1) + 1)
		}
		q := db.Query{
			Name:     "main",
			QueryRaw: "INSERT INTO banners (title, text, url, feature_id, tag_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		}
		args := []interface{}{title, text, url, feature, tags}
		db := app.ServiceProvider.BDClient(ctx)
		db.DB().QueryRowContext(ctx, q, args...)
	}
}
