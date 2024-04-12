package datamodel

import (
	"database/sql"
	"time"
)

type Banner struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Text      string `db:"text"`
	Url       string `db:"url"`
	IsActive  bool   `db:"active"`
	IsValid   bool
	Feature   int          `db:"feature_id"`
	Tag       []int        `db:"tag_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserBanner struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Url    string `db:"url"`
	Active bool   `db:"use_active-version"`
}
