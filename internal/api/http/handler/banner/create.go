package banner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//h.service.Create(r.Context(),)
	fmt.Println()
}
