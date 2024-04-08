package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	// GetUserBanner()
	// GetAllBanners()
	// CreateBanner()
	// UpdateBanner()
	// DeleteBanner()
	http.Handler
	AuthHandler
	BannerHandler
}

type AuthHandler interface {
}

type BannerHandler interface {
	GetUserBanner(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllBanners(http.ResponseWriter, *http.Request, httprouter.Params)
	CreateBanner(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteBanner(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateBanner(http.ResponseWriter, *http.Request, httprouter.Params)
}
