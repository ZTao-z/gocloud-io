package service

import (
	"html/template"
	"net/http"
)

func index2Info(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	usr := req.Form["username"][0]
	pwd := req.Form["password"][0]
	mail := req.Form["mail"][0]

	page := template.Must(template.ParseFiles("templates/info.html"))

	page.Execute(w, map[string]string{
		"usr":  usr,
		"pwd":  pwd,
		"mail": mail,
	})
}
