package handler

import (
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/api/http/handler/banner"
	"github.com/drizzleent/banners/internal/api/http/handler/control"
	"github.com/drizzleent/banners/internal/service"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	api.BannerHandler
	api.ControlHandler
	http.Handler
}

func NewHandler(bservice service.BannerService) *handler {
	h := &handler{}
	bannerHandler := banner.New(bservice)
	controlHandler := control.New(bannerHandler)
	r := initRoutes(bannerHandler, controlHandler)
	h.Handler = r
	return h
}

func initRoutes(bh api.BannerHandler, ah api.ControlHandler) *httprouter.Router {
	r := httprouter.New()
	r.GET("/user_banner", bh.GetUserBanner)
	r.GET("/banner", ah.Control)
	r.POST("/banner", ah.Control)
	r.PATCH("/banner/:id", ah.Control)
	r.DELETE("/banner/:id", ah.Control)
	return r
}
