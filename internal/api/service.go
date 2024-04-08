package api

import "net/http"

type Handler interface {
	// GetUserBanner()
	// GetAllBanners()
	// CreateBanner()
	// UpdateBanner()
	// DeleteBanner()
	http.Handler
}
