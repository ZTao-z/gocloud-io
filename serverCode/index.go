package service

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, req *http.Request) {
	page := template.Must(template.ParseFiles("templates/index.html"))
	page.Execute(w, nil)
}
