package app

import (
	"context"
	"log"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/api/http/handler"
	"github.com/drizzleent/banners/internal/config"
	"github.com/drizzleent/banners/internal/config/env"
)

type serviceProvider struct {
	pgCfg   config.PGConfig
	httpCfg config.HTTPConfig

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

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler()
	}

	return s.handler
}
