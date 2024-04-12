package banner

import "github.com/drizzleent/banners/internal/service"

const (
	idQuery             = "id"
	tagIdQuery          = "tag_id"
	featureIdQuery      = "feature_id"
	useLastVersionQuery = "use_last_version"
	limitQuery          = "limit"
	offsetQuery         = "offset"
	tokenQuery          = "token"
)

type bannerHandler struct {
	service service.BannerService
}

func New(srv service.BannerService) *bannerHandler {
	return &bannerHandler{
		service: srv,
	}
}
