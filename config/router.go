package config

import (
	"github.com/julienschmidt/httprouter"

	"github.com/linggaaskaedo/go-playground-wire/controller"
	"github.com/linggaaskaedo/go-playground-wire/exception"
)

func NewRouter(newsController controller.NewsController) *httprouter.Router {
	router := httprouter.New()

	// router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/news/:newsID", newsController.FindNewsByID)
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
