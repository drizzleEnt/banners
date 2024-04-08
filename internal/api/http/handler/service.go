package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	http.Handler
}

func NewHandler() *handler {
	h := &handler{}
	r := initRoutes(h)
	h.Handler = r
	return h
}

func initRoutes(h *handler) *httprouter.Router {
	r := httprouter.New()

	r.GET("/user_banner", h.GetUserBanner)
	r.GET("/banner", h.GetAllBanners)
	r.POST("/banner", h.CreateBanner)
	r.PATCH("/banner/:id", h.UpdateBanner)
	r.DELETE("/banner/:id", h.DeleteBanner)
	return r
}
