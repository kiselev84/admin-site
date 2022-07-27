package controller

import (
	"encoding/json"
	"strings"

	"html/template"
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

		checkIp := c.usecase.GetAll()

		//указываем путь к файлу с шаблоном
		tmpl, err := template.ParseFiles("../internal/ui/templates/check_ip.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//исполняем именованный шаблон "check_ip.html", передавая туда массив со списком пользователей
		err = tmpl.ExecuteTemplate(w, "check_ip", checkIp)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
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

		_, err = c.usecase.AddNewIp(newIp)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil)
			return
		}

		http.Redirect(w, r, "/ip", 301)
	} else {
		ts, err := template.ParseFiles(filesCheck...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, "../internal/ui/templates/create.html")
	}

}

// Обработчик edit ip-Ceck net.
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

	}
	w.WriteHeader(http.StatusBadRequest)

}

// возвращаем пользователю страницу для редактирования объекта ip-Ceck net.
func (c *Controller) EditPageCheckIp(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles(filesCheck...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}

	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/edit/")

		ipCheck, err := c.usecase.GetIpCheckId(id)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		} else {
			tmpl, _ := template.ParseFiles("../internal/ui/templates/edit_check_ip.html")
			tmpl.Execute(w, ipCheck)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик удаление ip-check из хранилища.
func (c *Controller) DeleteIpCheck(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		id := strings.TrimPrefix(r.URL.Path, "/delete/")

		err := c.usecase.DeleteIpCheck(id)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, "/ip", 301)
	}
	w.WriteHeader(http.StatusBadRequest)

}
