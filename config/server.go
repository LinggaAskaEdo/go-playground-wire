package config

import (
	"fmt"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire/middleware"
	"github.com/linggaaskaedo/go-playground-wire/model/common"
)

type ServerOptions struct {
	AuthMiddleware *middleware.AuthMiddleware
	Config         common.Configuration
}

func NewServer(opts *ServerOptions) *http.Server {
	address := fmt.Sprintf(
		"%s:%s",
		opts.Config.Server.ServerHost,
		opts.Config.Server.ServerPort,
	)

	return &http.Server{
		Addr:    address,
		Handler: opts.AuthMiddleware,
	}
}
