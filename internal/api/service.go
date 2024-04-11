package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	http.Handler
	ControlHandler
	BannerHandler
	Control(http.ResponseWriter, *http.Request, httprouter.Params)
}

type BannerHandler interface {
	GetUserBanner(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllBanners(http.ResponseWriter, *http.Request, httprouter.Params)
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	Delete(http.ResponseWriter, *http.Request, httprouter.Params)
	Update(http.ResponseWriter, *http.Request, httprouter.Params)
}

type ControlHandler interface {
	Control(http.ResponseWriter, *http.Request, httprouter.Params)
}
