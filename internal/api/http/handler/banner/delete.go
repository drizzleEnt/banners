package banner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *bannerHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("delete banner with id:", ps.ByName("id"))
}
