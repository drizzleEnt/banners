package banner

import (
	"context"

	"github.com/drizzleent/banners/internal/model"
)

func (s *bannerService) GetAllBanners(ctx context.Context, specs *model.Specs) (*model.Banner, error) {
	res, err := s.repo.GetUserBanner(ctx, specs)

	if err != nil {
		return nil, err
	}
	return res, nil
}
