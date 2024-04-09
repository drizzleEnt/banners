package converter

import (
	"strconv"

	"github.com/drizzleent/banners/internal/model"
)

func FromReqToService(feature string, tag string, versionStr string) (*model.Specs, error) {
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