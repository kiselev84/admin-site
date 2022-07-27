package controller

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Обработчик для отображения содержимого log check net.
func (c *Controller) GetLogCheckNet(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ts, err := template.ParseFiles(filesCheck...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		log := c.usecase.GetLogCheckNet()

		//указываем путь к файлу с шаблоном
		tmpl, err := template.ParseFiles("../internal/ui/templates/log_check_ip.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//исполняем именованный шаблон "check_ip.html", передавая туда массив со списком пользователей
		err = tmpl.ExecuteTemplate(w, "log_check_ip", log)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusOK)

		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения содержимого log check net city.
func (c *Controller) GetLogCheckNetCity(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		ts, err := template.ParseFiles(filesCheck...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		city := strings.TrimPrefix(r.URL.Path, "/check_net/")
		log := c.usecase.GetLogCheckNetCity(city)

		//указываем путь к файлу с шаблоном
		tmpl, err := template.ParseFiles("../internal/ui/templates/log_check_ip_city.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//исполняем именованный шаблон, передавая туда массив со списком пользователей
		err = tmpl.ExecuteTemplate(w, "log_check_ip_city", log)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
