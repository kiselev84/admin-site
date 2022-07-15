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
	router.HandleFunc("/addip", controller.AddNewIpForm)
	router.Get("/ip", controller.GetAll)

	// Роуты Тестового полигона (Home page)
	router.Get("/", All)
	router.Get("/krdc", KRDC)
	router.Get("/mikrotik", Mikrotik)
	router.Get("/check_net", controller.GetLogCheckNet)
	router.Get("/ssh", controller.GetLogSsh)

	// Роуты проверки сети по офисам (Сheck page)
	router.Get("/check_net/{city}", controller.GetLogCheckNetCity)

	//Подключаем css к controller
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("../internal/ui/static/")})
	router.Handle("/static", http.NotFoundHandler())
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

}
