package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"project/test_site/internal/usecase"
)

func Build(router *chi.Mux, usecase usecase.Usecase) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	controller := NewController(usecase)

	// Роуты для проверяемых ip
	router.Post("/create", controller.AddNewIp)
	router.Get("/ip", controller.GetAll)

	//// Роуты Тестового полигона
	router.Get("/", All)
	router.Get("/KRDC", KRDC)
	router.Get("/Mikrotik", Mikrotik)
	router.Get("/Check", Check)

	//Подключаем css к controller
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("../internal/ui/static/")})
	router.Handle("/static", http.NotFoundHandler())
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

}
