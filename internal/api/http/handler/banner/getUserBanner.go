package banner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drizzleent/banners/internal/api"
	"github.com/drizzleent/banners/internal/converter"
	"github.com/drizzleent/banners/internal/interceptor"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *bannerHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpStatus, token, err := interceptor.CheckToken(r.Header.Get(tokenQuery))
	if err != nil {
		api.NewErrorResponse(w, httpStatus, err.Error())
		return
	}

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

	if _, err = interceptor.CheckAdminToken(token); err != nil {
		logrus.Print("Получение баннера для пользователя")
		w.Write([]byte("all banners are disabled according to the specified parameters"))
	}

	res, err := json.Marshal(banner)
	if err != nil {
		api.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Print("Получение баннера для пользователя")
	w.Write(res)
}
