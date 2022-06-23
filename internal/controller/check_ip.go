package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"project/test_site/internal/entity"
	"project/test_site/internal/usecase"
)

type Config struct {
}

type Controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func buildResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}

// Обработчик дабовления нового ip для проверки.
func (c *Controller) AddNewIp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := &entity.Ipcheck{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}
		id, err := c.usecase.AddNewIp(user)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}
		result := map[string]int{"id": int(id)}
		response, err := json.Marshal(result)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}
		buildResponse(w, http.StatusCreated, response)
		w.Write([]byte("\n"))
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения проверяемых ip.
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
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

		for _, user := range c.usecase.GetAll() {
			response += fmt.Sprintf("-     id:%v %v %s %s %v<br/>", user.Id, user.Ip, user.Office, user.City, user.Server)
		}

		_, err = io.WriteString(w, `<html><head><title>Проверка веб-службы</title></head><body><p>&nbsp;</p><h1 style="text-align: left;"><span style="color: #339966;"><strong>
		  Проверяемые IP:</strong></span></h1><div></div></body></html>`)
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
