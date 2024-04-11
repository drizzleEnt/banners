package converter

import (
	"strconv"

	"github.com/drizzleent/banners/internal/model"
)

func FromReqToUser(feature string, tag string, versionStr string) (*model.Specs, error) {
	feature_id, err := strconv.Atoi(feature)
	if err != nil {
		return nil, err
	}

	tag_id, err := strconv.Atoi(tag)
	if err != nil {
		return nil, err
	}

	version := versionStr == "true"

	return &model.Specs{
		Feature: feature_id,
		Tag:     tag_id,
		Version: version,
	}, nil
}

func FromReqToAdmin(feature string, tag string, limitQuery string, offsetQuery string) (*model.Specs, error) {
	feature_id, err := strconv.Atoi(feature)
	if err != nil {
		return nil, err
	}

	tag_id, err := strconv.Atoi(tag)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(feature)
	if err != nil {
		limit = 0
	}

	offset, err := strconv.Atoi(tag)
	if err != nil {
		offset = 0
	}

	return &model.Specs{
		Feature: feature_id,
		Tag:     tag_id,
		Limit:   limit,
		Offset:  offset,
		Version: false,
	}, nil
}

func FromReqToBanner() *model.Banner {

	return &model.Banner{
		Title:  "",
		Text:   "",
		Url:    "",
		Active: false,
	}
}
