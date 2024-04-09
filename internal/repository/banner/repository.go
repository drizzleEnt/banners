package banner

import (
	"context"
	"fmt"

	"github.com/drizzleent/banners/internal/model"
	"github.com/drizzleent/banners/internal/repository"
	"github.com/drizzleent/banners/internal/repository/conventer"
	datamodel "github.com/drizzleent/banners/internal/repository/data_model"
	"github.com/drizzleent/banners/pkg/client/db"
)

const (
	bannerTable   = "banners"
	idColumn      = "id"
	titleColumn   = "title"
	textColumn    = "text"
	urlColumn     = "url"
	featureColumn = "feature_id"
	tagColumn     = "tag_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUserBanner(ctx context.Context, specs *model.Specs) (*model.Banner, error) {
	var banner datamodel.Banner

	query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = $1 AND $2 = ANY(%s)", titleColumn, textColumn, urlColumn, bannerTable, featureColumn, tagColumn)

	q := db.Query{
		Name:     "repository.GetUserBanner",
		QueryRaw: query,
	}

	args := []interface{}{specs.Feature, specs.Tag}

	err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&banner.Title, &banner.Text, &banner.Url)
	return conventer.ToModelFromRepo(banner), err
}
