package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("web server directory error")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/unknown", errorNotice).Methods("GET")
	mx.HandleFunc("/getInfo", getInfo).Methods("GET")
	mx.HandleFunc("/", indexPage).Methods("GET")
	mx.HandleFunc("/", index2Info).Methods("POST")
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}
