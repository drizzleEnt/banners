package banner

import (
	"context"
	"fmt"

	"github.com/drizzleent/banners/internal/model"
)

func (s *bannerService) Create(ctx context.Context, model *model.Banner) (int64, error) {
	fmt.Println("service.Create")

	id, err := s.repo.Create(ctx, model)
	if err != nil {
		return 0, err
	}

	return id, nil
}
