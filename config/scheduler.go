package config

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/linggaaskaedo/go-playground-wire/service"
)

func NewScheduler(newsService service.NewsService) (gocron.Scheduler, error) {
	ctx := context.Background()

	s, err := gocron.NewScheduler()
	if err != nil {
		return s, err
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			newsService.ExtractNews,
			ctx,
		),
		// gocron.OneTimeJob(
		// 	gocron.OneTimeJobStartImmediately(),
		// ),
		// gocron.NewTask(
		// 	newsService.ExtractNews,
		// 	ctx,
		// ),
	)
	if err != nil {
		return s, err
	}

	// fmt.Println(j1.ID())

	// s.Start()

	return s, nil
}
