package app

import (
	"context"
	"log"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/api/http/handler"
	"github.com/drizzleent/banners/internal/api/http/handler/auth"
	"github.com/drizzleent/banners/internal/config"
	"github.com/drizzleent/banners/internal/config/env"
	"github.com/drizzleent/banners/internal/service"
	"github.com/drizzleent/banners/internal/service/banner"
)

type serviceProvider struct {
	pgCfg   config.PGConfig
	httpCfg config.HTTPConfig

	bannerService service.BannerService
	authService   service.AuthService

	handler api.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if nil == s.pgCfg {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to load pg config %s", err.Error())
		}

		s.pgCfg = cfg
	}

	return s.pgCfg
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if nil == s.httpCfg {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config %s", err.Error())
		}
		s.httpCfg = cfg
	}

	return s.httpCfg
}

func (s *serviceProvider) BannerService(ctx context.Context) service.BannerService {
	if nil == s.bannerService {
		s.bannerService = banner.New()
	}
	return s.bannerService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if nil == s.authService {
		s.authService = auth.New()
	}
	return s.authService
}

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler(s.BannerService(ctx), s.AuthService(ctx))
	}

	return s.handler
}
