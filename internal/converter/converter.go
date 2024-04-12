package converter

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

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

func FromReqToBanner(r *http.Request) (*model.Banner, int, error) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	var banner model.Banner

	err = json.Unmarshal(body, &banner)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if strings.Contains(string(body), "is_active") {
		banner.IsValid = false
	}

	return &banner, http.StatusOK, nil
}

func FromReqToUpdateBanner(r *http.Request) (*model.Banner, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	var banner model.Banner

	err = json.Unmarshal(body, &banner)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &banner, http.StatusOK, nil
}
