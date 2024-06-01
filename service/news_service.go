package service

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire/model/web"
)

type NewsService interface {
	ExtractNews(ctx context.Context)
	FindById(ctx context.Context, categoryId int) web.NewsResponse
}
