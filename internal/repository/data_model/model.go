package datamodel

import (
	"database/sql"
	"time"
)

type Banner struct {
	Title     string       `db:"title"`
	Text      string       `db:"text"`
	Url       string       `db:"url"`
	Active    bool         `db:"active"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserBanner struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Url    string `db:"url"`
	Active bool   `db:"use_active-version"`
}
