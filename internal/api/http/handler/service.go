package handler

import (
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/api/http/handler/auth"
	"github.com/drizzleent/banners/internal/api/http/handler/banner"
	"github.com/drizzleent/banners/internal/service"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	api.BannerHandler
	api.AuthHandler
	http.Handler
}

func NewHandler(bservice service.BannerService, aservice service.AuthService) *handler {
	h := &handler{}
	bannerHandler := banner.New(bservice)
	authHandler := auth.New()
	r := initRoutes(bannerHandler, authHandler)
	h.Handler = r
	return h
}

func initRoutes(bh api.BannerHandler, ah api.AuthHandler) *httprouter.Router {
	r := httprouter.New()

	r.GET("/user_banner", bh.GetUserBanner)
	r.GET("/banner", bh.GetAllBanners)
	r.POST("/banner", bh.CreateBanner)
	r.PATCH("/banner/:id", bh.UpdateBanner)
	r.DELETE("/banner/:id", bh.DeleteBanner)
	return r
}
