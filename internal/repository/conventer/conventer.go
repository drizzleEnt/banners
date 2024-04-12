package conventer

import (
	"github.com/drizzleent/banners/internal/model"
	datamodel "github.com/drizzleent/banners/internal/repository/data_model"
)

func ToModelFromRepo(banner datamodel.Banner) *model.Banner {
	return &model.Banner{
		Title:     banner.Title,
		Text:      banner.Text,
		Url:       banner.Url,
		IsActive:  banner.IsActive,
		IsValid:   banner.IsValid,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}

func ToUserModelFromRepo(banner datamodel.Banner) *model.UserBanner {
	return &model.UserBanner{
		Title: banner.Title,
		Text:  banner.Title,
		Url:   banner.Url,
	}
}

func ToUpdate(old datamodel.Banner, req *model.Banner) *model.Banner {

	if req.Title == "" {
		req.Title = old.Title
	}

	if req.Text == "" {
		req.Text = old.Text
	}

	if req.Url == "" {
		req.Url = old.Url
	}

	if req.Feature == 0 {
		req.Feature = old.Feature
	}

	if len(req.Tag) == 0 {
		req.Tag = old.Tag
	}

	if req.IsValid {

	}

	return req
}
