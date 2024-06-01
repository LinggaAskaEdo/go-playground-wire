package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/linggaaskaedo/go-playground-wire/helper"
	"github.com/linggaaskaedo/go-playground-wire/model/domain"
)

type NewsRepositoryImpl struct {
}

func NewNewsRepository() *NewsRepositoryImpl {
	return &NewsRepositoryImpl{}
}

func (repository *NewsRepositoryImpl) GetNewsByUrl(ctx context.Context, tx *sql.Tx, url string) (bool, error) {
	var isExist bool

	// SQL := "SELECT IF(COUNT(1), 'true', 'false') FROM news_article WHERE url = ?"
	err := tx.QueryRowContext(ctx, GetNewsByUrl, url).Scan(&isExist)
	helper.LogIfError(err)

	return isExist, nil
}

func (repository *NewsRepositoryImpl) CreateNews(ctx context.Context, tx *sql.Tx, v domain.NewsArticle) (domain.NewsArticle, error) {
	_, err := tx.ExecContext(ctx, CreateNews, v.Title, v.URL, v.Content, v.Summary, v.ArticleTS, v.PublishedDate, v.Inserted, v.Updated)
	helper.LogIfError(err)

	// id, err := result.LastInsertId()
	// helper.PanicIfError(err)

	// v.ID = int(id)

	return v, nil
}

func (repository *NewsRepositoryImpl) FindNewsByID(ctx context.Context, tx *sql.Tx, categoryId int) (domain.News, error) {
	rows, err := tx.QueryContext(ctx, FindNewsByID, categoryId)
	helper.LogIfError(err)
	defer rows.Close()

	news := domain.News{}
	if rows.Next() {
		err := rows.Scan(&news.Id, &news.Title, &news.Url, &news.Content)
		helper.LogIfError(err)

		return news, nil
	} else {
		return news, errors.New("News is not found")
	}
}
