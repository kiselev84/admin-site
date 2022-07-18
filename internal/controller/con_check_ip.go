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
	"strconv"
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

		http.ServeFile(w, r, "../internal/ui/templates/button_add_ip.html")
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

// Обработчик дабовления нового ip для проверки HTML form.
func (c *Controller) AddNewIpForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		newIp := &entity.Ipcheck{}

		newIp.Office = r.FormValue("office")
		newIp.Ip = r.FormValue("ip")
		newIp.City = r.FormValue("city")
		newIp.Server = r.FormValue("device")

		if err != nil {
			log.Println(err)
		}
		_, err = c.usecase.AddNewIp(newIp)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}

		http.Redirect(w, r, "/ip", 301)
	} else {
		_, err := io.WriteString(w, `<html><head><title>Проверка веб-службы</title>
			</head><body><p>&nbsp;</p><h1 style="text-align: left;">
			<span style="color: #339966;"><strong>Добавление IP:</strong></span>
			</h1><div></div></body>
			</html>`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, "../internal/ui/templates/create.html")
	}

}

// Обработчик edit-ip Ceck net.
func (c *Controller) EditCheckIp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		editIp := &entity.Ipcheck{}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
		}
		editIp.Id = uint8(id)
		editIp.Office = r.FormValue("office")
		editIp.Ip = r.FormValue("ip")
		editIp.City = r.FormValue("city")
		editIp.Server = r.FormValue("device")

		if err != nil {
			log.Println(err)
		}
		err = c.usecase.EditCheckIp(editIp)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}

		http.Redirect(w, r, "/ip", 301)
	} else {
		_, err := io.WriteString(w, `<html><head><title>Проверка веб-службы</title>
			</head><body><p>&nbsp;</p><h1 style="text-align: left;">
			<span style="color: #339966;"><strong>Редактирование IP:</strong></span>
			</h1><div></div></body>
			</html>`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, "../internal/ui/templates/edit_check_ip.html")
	}

}
