package banner

import (
	"context"

	"github.com/drizzleent/banners/internal/model"
)

func (s *bannerService) Update(ctx context.Context, id int64, model *model.Banner) error {
	err := s.repo.Update(ctx, id, model)
	if err != nil {
		return err
	}
	return nil
}
