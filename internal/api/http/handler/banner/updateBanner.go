package banner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) UpdateBanner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("update banner with id:", ps.ByName("id"))
}
