package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire/model/common"
)

type CacheOptions struct {
	Config common.Configuration
}

func NewCache(opts *CacheOptions) (*redis.Client, error) {
	ctx := context.Background()

	cacheURI := fmt.Sprintf(
		"%s:%s",
		opts.Config.Cache.CacheHost,
		opts.Config.Cache.CachePort,
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     cacheURI,
		Password: opts.Config.Database.DBPassword,
	})

	status := rdb.Ping(ctx)
	if status.Val() != "PONG" {
		return rdb, errors.New(status.String())
	}

	return rdb, nil
}
