package control

import (
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/interceptor"
	"github.com/julienschmidt/httprouter"
)

type controlHandler struct {
	api.BannerHandler
}

func New(bh api.BannerHandler) *controlHandler {
	return &controlHandler{
		BannerHandler: bh,
	}
}

func (h *controlHandler) Control(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httpStatus, err := interceptor.CheckAdminToken(r.Header.Get("token"))
	if err != nil {
		api.NewErrorResponse(w, httpStatus, err.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.BannerHandler.GetAllBanners(w, r, p)
	case http.MethodPost:
		h.BannerHandler.Create(w, r, p)
	case http.MethodPatch:
		h.BannerHandler.Update(w, r, p)
	case http.MethodDelete:
		h.BannerHandler.Delete(w, r, p)
	}
}
