package api

type Handler interface {
	GetUserBanner()
	GetAllBanners()
	CreateBanner()
	UpdateBanner()
	DeleteBanner()
}
