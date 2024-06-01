package repository

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-playground-wire/model/domain"
)

type NewsRepository interface {
	GetNewsByUrl(ctx context.Context, tx *sql.Tx, url string) (bool, error)
	CreateNews(ctx context.Context, tx *sql.Tx, v domain.NewsArticle) (domain.NewsArticle, error)
	FindNewsByID(ctx context.Context, tx *sql.Tx, categoryId int) (domain.News, error)
}
