package config

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/linggaaskaedo/go-playground-wire/model/common"
	"github.com/linggaaskaedo/go-playground-wire/service"
)

type SchedulerOptions struct {
	Config common.Configuration
}

func NewScheduler(newsService service.NewsService, opts *SchedulerOptions) (gocron.Scheduler, error) {
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
			opts.Config.RSS.RssURL,
		),
	)
	if err != nil {
		return s, err
	}

	return s, nil
}
