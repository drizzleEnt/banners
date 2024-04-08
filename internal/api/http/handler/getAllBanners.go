package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) GetAllBanners(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("get all banners")
}
