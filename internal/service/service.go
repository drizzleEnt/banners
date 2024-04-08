package service

import "context"

type BannerService interface {
	GetAllBanners(context.Context) error
}

type AuthService interface {
}
