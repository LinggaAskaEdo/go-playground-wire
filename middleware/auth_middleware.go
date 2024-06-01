package middleware

import (
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire/helper"
	"github.com/linggaaskaedo/go-playground-wire/model/common"
	"github.com/linggaaskaedo/go-playground-wire/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
	Config  common.Configuration
}

func NewAuthMiddleware(handler http.Handler, Config common.Configuration) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
		Config:  Config,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if middleware.Config.Token.Auth == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
