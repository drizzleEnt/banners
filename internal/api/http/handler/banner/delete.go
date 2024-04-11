package banner

import (
	"net/http"
	"strconv"

	"github.com/drizzleent/banners/internal/api"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *bannerHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(r.Context(), int64(id))

	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Баннер успешно удален"))

	logrus.Printf("delete banner with id:%v\n", id)
}
