package service

import (
	"net/http"
)

func getInfo(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome to cloudgo-io"))
}
