package repository

import (
	"context"

	"github.com/drizzleent/banners/internal/model"
)

type Repository interface {
	BannerRepository
}

type BannerRepository interface {
	GetUserBanner(context.Context, *model.Specs) (*model.Banner, error)
}
