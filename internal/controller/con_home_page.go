package controller

import (
	"fmt"
	"html/template"
	"io"
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

}

// Обработчик для отображения содержимого KRDC.
func KRDC(w http.ResponseWriter, r *http.Request) {

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

// Обработчик для отображения содержимого Mikrotik.
func Mikrotik(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles(filesMikrotik...)
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

// Обработчик для отображения содержимого log ssh.
func (c *Controller) GetLogSsh(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "GET" {
		response := ""

		for _, user := range c.usecase.GetLogSsh() {
			response += fmt.Sprintf("-     id:%v %v %v %v<br/>", user.Id, user.Time, user.Ip, user.Text)
		}

		_, err = io.WriteString(w, `<html><head><title>Проверка веб-службы</title></head><body><p>&nbsp;</p><h1 style="text-align: left;"><span style="color: #339966;"><strong>
		  Лог проверки SSH:</strong></span></h1><div></div></body></html>`)
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
