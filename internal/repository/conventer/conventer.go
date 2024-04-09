package conventer

import (
	"github.com/drizzleent/banners/internal/model"
	datamodel "github.com/drizzleent/banners/internal/repository/data_model"
)

func ToModelFromRepo(banner datamodel.Banner) *model.Banner {
	return &model.Banner{
		Title: banner.Title,
		Text:  banner.Text,
		Url:   banner.Url,
	}
}
