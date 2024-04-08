package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) UpdateBanner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("update banner with id: ", ps.ByName("id"))
}
