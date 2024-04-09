package datamodel

type Banner struct {
	Title string `db:"title"`
	Text  string `db:"text"`
	Url   string `db:"url"`
}
