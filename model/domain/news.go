package domain

import "database/sql"

type News struct {
	Id      int
	Title   string
	Url     string
	Content string
}

type NewsArticle struct {
	Title         string       `db:"title"`
	URL           string       `db:"url"`
	Content       string       `db:"content"`
	Summary       string       `db:"summary"`
	ArticleTS     int64        `db:"article_ts"`
	PublishedDate sql.NullTime `db:"published_date"`
	Inserted      sql.NullTime `db:"inserted"`
	Updated       sql.NullTime `db:"updated"`
}
