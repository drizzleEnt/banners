package repository

import (
	"context"

	"github.com/drizzleent/banners/internal/model"
)

type Repository interface {
	BannerRepository
}

type BannerRepository interface {
	GetUserBanner(context.Context, *model.Specs) (*model.UserBanner, error)
	GetAllBanners(context.Context, *model.Specs) (*model.Banner, error)
	Create(context.Context, *model.Banner) (int64, error)
	Update(context.Context, *model.Banner) error
	Delete(context.Context, int64) error
}
