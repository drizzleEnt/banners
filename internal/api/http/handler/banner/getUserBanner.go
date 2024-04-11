package banner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/drizzleent/banners/internal/interceptor"
	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := interceptor.CheckToken(r.Header.Get(tokenQuery))
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// feature_id := r.URL.Query().Get(featureIdQuery)
	// if len(feature_id) == 0 {
	// 	api.NewErrorResponse(w, http.StatusBadRequest, "Некорректные данные feature id requared")
	// 	return
	// }

	// tag_id := r.URL.Query().Get(tagIdQuery)
	// if len(tag_id) == 0 {
	// 	api.NewErrorResponse(w, http.StatusBadRequest, "Некорректные данные tag id requared")
	// 	return
	// }
	// useLastVersion := r.URL.Query().Get(useLastVersionQuery)

	specs, err := converter.FromReqToUser(
		r.URL.Query().Get(featureIdQuery),
		r.URL.Query().Get(tagIdQuery),
		r.URL.Query().Get(useLastVersionQuery))

	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Некорректные данные %s", err.Error()))
		return
	}
	banner, err := h.service.GetUserBanner(r.Context(), specs)

	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(banner)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write(res)
}
