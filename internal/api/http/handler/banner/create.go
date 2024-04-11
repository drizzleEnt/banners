package banner

import (
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	banner, httpStatus, err := converter.FromReqToBanner(r)
	if err != nil {
		api.NewErrorResponse(w, httpStatus, err.Error())
		return
	}
	id, err := h.service.Create(r.Context(), banner)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(id)
}
