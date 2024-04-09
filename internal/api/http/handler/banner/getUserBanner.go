package banner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	feature_id := r.URL.Query().Get(featureIdQuery)
	if len(feature_id) == 0 {
		api.NewErrorResponse(w, http.StatusBadRequest, "feature id requared")
		return
	}

	tag_id := r.URL.Query().Get(tagIdQuery)
	if len(tag_id) == 0 {
		api.NewErrorResponse(w, http.StatusBadRequest, "tag id requared")
		return
	}
	useLastVersion := r.URL.Query().Get(useLastVersionQuery)

	token := r.Header.Get(tokenQuery)
	if len(token) == 0 {
		api.NewErrorResponse(w, http.StatusBadRequest, "token requared")
		return
	}
	fmt.Println(feature_id, tag_id, token)

	specs, err := converter.FromReqToService(feature_id, tag_id, useLastVersion)
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("incorrect data %s", err.Error()))
		return
	}
	banner, err := h.service.GetAllBanners(r.Context(), specs)

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
	fmt.Println(res)
}
