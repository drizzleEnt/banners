package banner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) GetAllBanners(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	specs, err := converter.FromReqToAdmin(
		r.URL.Query().Get(featureIdQuery),
		r.URL.Query().Get(tagIdQuery),
		r.URL.Query().Get(limitQuery),
		r.URL.Query().Get(offsetQuery))
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Некорректные данные %s", err.Error()))
	}
	banners, err := h.service.GetAllBanners(r.Context(), specs)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	res, err := json.Marshal(banners)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write(res)
}
