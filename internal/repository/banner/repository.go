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
	activeColumn  = "active"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUserBanner(ctx context.Context, specs *model.Specs) (*model.UserBanner, error) {
	var banner datamodel.Banner

	query := fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = $1 AND $2 = ANY(%s)", titleColumn, textColumn, urlColumn, activeColumn, bannerTable, featureColumn, tagColumn)

	q := db.Query{
		Name:     "repository.GetUserBanner",
		QueryRaw: query,
	}

	args := []interface{}{specs.Feature, specs.Tag}

	err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&banner.Title, &banner.Text, &banner.Url, &banner.Active)
	return conventer.ToUserModelFromRepo(banner), err
}

func (r *repo) GetAllBanners(ctx context.Context, specs *model.Specs) (*model.Banner, error) {
	return nil, nil
}

func (r *repo) Create(ctx context.Context, banner *model.Banner) (int64, error) {
	fmt.Println("repo.Create")
	fmt.Println(banner.Tag)
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		bannerTable, titleColumn, textColumn, urlColumn, activeColumn, featureColumn, tagColumn)
	q := db.Query{
		Name:     "repository.Create",
		QueryRaw: query,
	}

	args := []interface{}{banner.Title, banner.Text, banner.Url, banner.Active, banner.Feature, banner.Tag}

	var id int64
	err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repo) Update(context.Context, *model.Banner) error {
	return nil
}

func (r *repo) Delete(context.Context, int64) error {
	return nil
}
