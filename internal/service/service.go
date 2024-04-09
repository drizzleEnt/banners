package service

import (
	"context"

	"github.com/drizzleent/banners/internal/model"
)

type BannerService interface {
	GetAllBanners(context.Context, *model.Specs) (*model.Banner, error)
}

type AuthService interface {
}
