package controller

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	filesHome = []string{
		"../internal/ui/html/home.page.tmpl",
		"../internal/ui/html/base.layout.tmpl",
		"../internal/ui/html/footer.partial.tmpl",
	}
	filesKRDC = []string{
		"../internal/ui/html/krdc.page.html",
		"../internal/ui/html/base.layout.tmpl",
		"../internal/ui/html/footer.partial.tmpl",
	}
	filesMikrotik = []string{
		"../internal/ui/html/mikrotik.page.html",
		"../internal/ui/html/base.layout.tmpl",
		"../internal/ui/html/footer.partial.tmpl",
	}
	filesCheck = []string{
		"../internal/ui/html/check.page.html",
		"../internal/ui/html/base.layout.tmpl",
		"../internal/ui/html/footer.partial.tmpl",
	}
	filesCheckForm = []string{
		"../internal/ui/html/check_Rostov.page.html",
		"../internal/ui/html/base.layout.tmpl",
		"../internal/ui/html/footer.partial.tmpl",
	}
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}

// Обработчик главной странице.
func All(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/check_net", 301)
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения содержимого KRDC.
func KRDC(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, err := template.ParseFiles(filesKRDC...)
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
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения содержимого Mikrotik.
func Mikrotik(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, err := template.ParseFiles(filesHome...)
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

		http.ServeFile(w, r, "../internal/ui/templates/mikrotik.html")
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Обработчик для отображения содержимого log ssh.
func (c *Controller) GetLogSsh(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, err := template.ParseFiles(filesHome...)
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

		logSsh := c.usecase.GetLogSsh()

		//указываем путь к файлу с шаблоном
		tmpl, err := template.ParseFiles("../internal/ui/templates/log_shh.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//исполняем именованный шаблон "log_ssh", передавая туда массив со списком пользователей
		err = tmpl.ExecuteTemplate(w, "log_ssh", logSsh)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
