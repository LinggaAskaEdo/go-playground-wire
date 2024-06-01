//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"

	"github.com/linggaaskaedo/go-playground-wire/config"
	"github.com/linggaaskaedo/go-playground-wire/controller"
	"github.com/linggaaskaedo/go-playground-wire/middleware"
	"github.com/linggaaskaedo/go-playground-wire/model/common"
	"github.com/linggaaskaedo/go-playground-wire/repository"
	"github.com/linggaaskaedo/go-playground-wire/service"
)

var newsSet = wire.NewSet(
	repository.NewNewsRepository,
	wire.Bind(new(repository.NewsRepository), new(*repository.NewsRepositoryImpl)),
	service.NewNewsService,
	wire.Bind(new(service.NewsService), new(*service.NewsServiceImpl)),
	controller.NewNewsController,
	wire.Bind(new(controller.NewsController), new(*controller.NewsControllerImpl)),
)

func InitializedServer(configuration common.Configuration) (*http.Server, error) {
	wire.Build(
		wire.NewSet(wire.Struct(new(config.DBOptions), "*"), config.NewDB),
		wire.NewSet(wire.Struct(new(config.CacheOptions), "*"), config.NewCache),
		config.NewValidator,
		newsSet,
		config.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		wire.NewSet(wire.Struct(new(config.ServerOptions), "*"), config.NewServer),
	)

	return nil, nil
}

func InitializedScheduler(configuration common.Configuration) (gocron.Scheduler, error) {
	wire.Build(
		wire.NewSet(wire.Struct(new(config.DBOptions), "*"), config.NewDB),
		wire.NewSet(wire.Struct(new(config.CacheOptions), "*"), config.NewCache),
		config.NewValidator,
		repository.NewNewsRepository,
		wire.Bind(new(repository.NewsRepository), new(*repository.NewsRepositoryImpl)),
		service.NewNewsService,
		wire.Bind(new(service.NewsService), new(*service.NewsServiceImpl)),
		wire.NewSet(wire.Struct(new(config.SchedulerOptions), "*"), config.NewScheduler),
		// config.NewScheduler,
	)

	return nil, nil
}
