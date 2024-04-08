package banner

import "github.com/drizzleent/banners/internal/service"

type bannerService struct {
}

func New() service.BannerService {
	return &bannerService{}
}
