package app

import (
	"context"
	"log"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/api/http/handler"
	authSrv "github.com/drizzleent/banners/internal/api/http/handler/auth"
	"github.com/drizzleent/banners/internal/config"
	"github.com/drizzleent/banners/internal/config/env"
	"github.com/drizzleent/banners/internal/repository"
	bannerRepo "github.com/drizzleent/banners/internal/repository/banner"
	"github.com/drizzleent/banners/internal/service"
	bannerSrv "github.com/drizzleent/banners/internal/service/banner"
	"github.com/drizzleent/banners/pkg/client/db"
	"github.com/drizzleent/banners/pkg/client/db/pg"
)

type serviceProvider struct {
	pgCfg   config.PGConfig
	httpCfg config.HTTPConfig

	db db.Client

	repository repository.Repository

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

func (s *serviceProvider) BDClient(ctx context.Context) db.Client {
	if nil == s.db {
		cl, err := pg.New(ctx, s.PGConfig().Address())
		if err != nil {
			log.Fatalf("Failed to create db client %s", err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("Failed to ping db %s", err.Error())
		}

		s.db = cl
	}

	return s.db
}

func (s *serviceProvider) Repository(ctx context.Context) repository.Repository {
	if nil == s.repository {
		s.repository = bannerRepo.NewRepository(s.BDClient(ctx))
	}

	return s.repository
}

func (s *serviceProvider) BannerService(ctx context.Context) service.BannerService {
	if nil == s.bannerService {
		s.bannerService = bannerSrv.New(s.Repository(ctx))
	}
	return s.bannerService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if nil == s.authService {
		s.authService = authSrv.New()
	}
	return s.authService
}

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler(s.BannerService(ctx), s.AuthService(ctx))
	}

	return s.handler
}
