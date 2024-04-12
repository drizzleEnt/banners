package banner

import (
	"net/http"
	"strconv"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *bannerHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName(idQuery))
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	banner, httpStatus, err := converter.FromReqToBanner(r)
	if err != nil {
		api.NewErrorResponse(w, httpStatus, err.Error())
		return
	}

	err = h.service.Update(r.Context(), int64(id), banner)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("Обновление содержимого баннера id: %v\n", id)
}
