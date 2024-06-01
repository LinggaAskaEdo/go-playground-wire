package helper

import (
	"github.com/linggaaskaedo/go-playground-wire/model/domain"
	"github.com/linggaaskaedo/go-playground-wire/model/web"
)

func ToNewsResponse(news domain.News) web.NewsResponse {
	return web.NewsResponse{
		Id:      news.Id,
		Title:   news.Title,
		Url:     news.Url,
		Content: news.Content,
	}
}

func ToNewsResponses(arrNews []domain.News) []web.NewsResponse {
	var newsResponses []web.NewsResponse

	for _, news := range arrNews {
		newsResponses = append(newsResponses, ToNewsResponse(news))
	}

	return newsResponses
}
