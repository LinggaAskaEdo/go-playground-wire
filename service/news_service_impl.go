package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire/exception"
	"github.com/linggaaskaedo/go-playground-wire/helper"
	"github.com/linggaaskaedo/go-playground-wire/model/web"
	"github.com/linggaaskaedo/go-playground-wire/repository"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	DB             *sql.DB
	Cache          *redis.Client
	Validate       *validator.Validate
}

func NewNewsService(newsRepository repository.NewsRepository, DB *sql.DB, Cache *redis.Client, validate *validator.Validate) *NewsServiceImpl {
	return &NewsServiceImpl{
		NewsRepository: newsRepository,
		DB:             DB,
		Cache:          Cache,
		Validate:       validate,
	}
}

func (service *NewsServiceImpl) ExtractNews(ctx context.Context) {
	log.Println("JOB RUNNING !!!")
}

func (service *NewsServiceImpl) FindById(ctx context.Context, categoryId int) web.NewsResponse {
	log.Println("SERVICE RUNNING !!!")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	news, err := service.NewsRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewsResponse(news)
}
