package model

import (
	"database/sql"
	"time"
)

type Banner struct {
	ID        int
	Title     string `json:"title,omitempty"`
	Text      string `json:"text,omitempty"`
	Url       string `json:"url,omitempty"`
	IsActive  bool   `json:"is_active,omitempty"`
	IsValid   bool   `json:"is_valid"`
	Feature   int    `json:"feature_id,omitempty"`
	Tag       []int  `json:"tag_id,omitempty"`
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserBanner struct {
	Title    string `json:"title,omitempty"`
	Text     string `json:"text,omitempty"`
	Url      string `json:"url,omitempty"`
	Feature  int    `json:"feature_id,omitempty"`
	Tag      []int  `json:"tag_id,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
	IsValid  bool
}

type Specs struct {
	Feature int
	Tag     int
	Limit   int
	Offset  int
	Version bool
}
