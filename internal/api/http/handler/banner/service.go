package banner

import "github.com/drizzleent/banners/internal/service"

type bannerHandler struct {
	service service.BannerService
}

func New(srv service.BannerService) *bannerHandler {
	return &bannerHandler{
		service: srv,
	}
}
