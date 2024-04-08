package banner

import (
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) GetAllBanners(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	feature_id := r.Header.Get("feature_id")
	if len(feature_id) == 0 {
		api.NewErrorResponse(w, http.StatusBadRequest, "feature id requared")
		return
	}
	tag_id := r.Header.Get("tag_id")
	if len(tag_id) == 0 {
		api.NewErrorResponse(w, http.StatusBadRequest, "tag id requared")
		return
	}

	fmt.Println(feature_id, tag_id)

	err := h.service.GetAllBanners(r.Context())
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Write([]byte(feature_id + tag_id))
}
