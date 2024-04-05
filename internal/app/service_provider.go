package app

import (
	"log"

	"github.com/drizzleent/banners/internal/config"
	"github.com/drizzleent/banners/internal/config/env"
)

type serviceProvider struct {
	pgCfg   config.PGConfig
	httpCfg config.HTTPConfig
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
