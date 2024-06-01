package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/linggaaskaedo/go-playground-wire/helper"
	"github.com/linggaaskaedo/go-playground-wire/model/web"
	"github.com/linggaaskaedo/go-playground-wire/service"
)

type NewsControllerImpl struct {
	NewsService service.NewsService
}

func NewNewsController(newsService service.NewsService) *NewsControllerImpl {
	return &NewsControllerImpl{
		NewsService: newsService,
	}
}

func (controller *NewsControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	newsID := params.ByName("newsID")
	id, err := strconv.Atoi(newsID)
	helper.PanicIfError(err)

	newsResponse := controller.NewsService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
