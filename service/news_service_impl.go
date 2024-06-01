package service

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire/helper"
	"github.com/linggaaskaedo/go-playground-wire/model/domain"
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

func (service *NewsServiceImpl) ExtractNews(ctx context.Context, url string) {
	log.Println("JOB RUNNING !!!")

	start := time.Now()

	counter := 0

	doc, err := xmlquery.LoadURL(url)
	if err != nil {
		helper.LogIfError(err)
	}

	v := domain.NewsArticle{}

	channel := xmlquery.Find(doc, "//item")

	for _, n := range channel {
		if n := n.SelectElement("title"); n != nil {
			title := n.InnerText()
			v.Title = title
		}

		if n := n.SelectElement("link"); n != nil {
			link := n.InnerText()
			v.URL = link

			docDetail, err := htmlquery.LoadURL(link)
			if err != nil {
				helper.LogIfError(err)
			}

			docDataDetail := htmlquery.FindOne(docDetail, "//div[@class = 'post-content clearfix']")
			strDocDataDetail := htmlquery.InnerText(docDataDetail)
			strDocDataDetail = strings.TrimSpace(strDocDataDetail)
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\t", "")
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\n", "")

			v.Content = strDocDataDetail
		}

		if n := n.SelectElement("pubDate"); n != nil {
			pubDate := n.InnerText()

			timePub, err := time.Parse(time.RFC1123Z, pubDate)
			if err != nil {
				helper.LogIfError(err)
			}

			v.PublishedDate = sql.NullTime{Time: timePub, Valid: true}

			timestamp := timePub.Unix()
			v.ArticleTS = int64(timestamp)
		}

		v.Inserted = sql.NullTime{Time: time.Now(), Valid: true}

		tx, err := service.DB.Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollback(tx)

		status, err := service.NewsRepository.GetNewsByUrl(ctx, tx, v.URL)
		if err != nil {
			helper.LogIfError(err)
		}

		log.Println("URL: ", v.URL, ", Status: ", status)

		if !status {
			_, err := service.NewsRepository.CreateNews(ctx, tx, v)
			if err != nil {
				log.Fatalln(err)
			}

			counter++
		}
	}

	duration := time.Since(start)

	log.Println(counter, " data added successfully in ", duration.Seconds(), " seconds")
}

func (service *NewsServiceImpl) FindNewsByID(ctx context.Context, categoryId int) web.NewsResponse {
	log.Println("SERVICE RUNNING !!!")

	tx, err := service.DB.Begin()
	helper.LogIfError(err)
	defer helper.CommitOrRollback(tx)

	news, err := service.NewsRepository.FindNewsByID(ctx, tx, categoryId)
	helper.LogIfError(err)

	return helper.ToNewsResponse(news)
}
