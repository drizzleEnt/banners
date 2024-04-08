package banner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("create banner")
}
