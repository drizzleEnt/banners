package banner

import (
	"github.com/drizzleent/banners/internal/repository"
	"github.com/drizzleent/banners/internal/service"
)

type bannerService struct {
	repo repository.Repository
}

func New(repo repository.Repository) service.BannerService {
	return &bannerService{
		repo: repo,
	}
}
