package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NewsController interface {
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
