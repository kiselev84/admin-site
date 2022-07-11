package controller

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Обработчик для отображения содержимого log check net.
func (c *Controller) GetLogCheckNet(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "GET" {
		response := ""

		for _, user := range c.usecase.GetLogCheckNet() {
			response += fmt.Sprintf("-     ID:%v %v %v %v %v %v %v<br/>", user.Id, user.Text, user.Time, user.City, user.Office, user.Server, user.Ip)
		}

		_, err = io.WriteString(w, `<html><head><title>Проверка веб-службы</title></head><body><p>&nbsp;</p><h1 style="text-align: left;"><span style="color: #339966;"><strong>
		  Лог проверки сети:</strong></span></h1><div></div></body></html>`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения содержимого log check net city.
func (c *Controller) GetLogCheckNetCity(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "GET" {
		response := ""

		city := strings.TrimPrefix(r.URL.Path, "/check_net/")
		for _, user := range c.usecase.GetLogCheckNetCity(string(city)) {
			response += fmt.Sprintf("-     ID:%v %v %v %v %v %v %v<br/>", user.Id, user.Text, user.Time, user.City, user.Office, user.Server, user.Ip)
		}

		_, err = io.WriteString(w, `<html><head><title>Проверка веб-службы</title></head><body><p>&nbsp;</p><h1 style="text-align: left;"><span style="color: #339966;"><strong>
		  Лог проверки сети по офису city: </strong></span></h1><div></div></body></html>`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
