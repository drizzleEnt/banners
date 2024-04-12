package banner

import (
	"context"
	"fmt"

	"github.com/drizzleent/banners/internal/model"
)

func (s *bannerService) GetUserBanner(ctx context.Context, specs *model.Specs) (*model.UserBanner, error) {
	res, err := s.repo.GetUserBanner(ctx, specs)

	if err != nil {
		return nil, err
	}

	if !res.IsActive {
		return nil, fmt.Errorf("banner is disable")
	}
	return res, nil
}
