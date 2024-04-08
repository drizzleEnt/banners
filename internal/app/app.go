package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/drizzleent/banners/internal/config"
	"github.com/drizzleent/banners/pkg/closer"
)

type App struct {
	serviceProvider *serviceProvider

	httpServer *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebs(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	err := a.runHttpServer()
	if err != nil {
		log.Fatalf("failed to run HTTP server: %s", err.Error())
	}
	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.httpServer = &http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        a.serviceProvider.Handler(ctx),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("HTTP server running on: %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
