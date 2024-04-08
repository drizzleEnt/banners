package auth

import "github.com/drizzleent/banners/internal/service"

type authService struct {
}

func New() service.AuthService {
	return &authService{}
}
