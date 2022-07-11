package app

import (
	"github.com/go-chi/chi"
	"net/http"
	"project/test_site/internal/controller"
	"project/test_site/internal/functions"
	"project/test_site/internal/repository"
	"project/test_site/internal/usecase"
)

func Run() {
	go functions.CheckShh()
	go functions.CheckNet()

	repository := repository.NewRepository()
	usecase := usecase.NewUsecase(repository)

	router := chi.NewRouter()
	controller.Build(router, usecase)
	http.ListenAndServe(":7777", router)

}
