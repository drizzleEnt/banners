package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) GetUserBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("get user banner")
}
