package service

import (
	"net/http"
)

// Error Page
func errorNotice(reqw http.ResponseWriter, req *http.Request) {
	http.Error(reqw, "501 Not Implemented!", 501)
}
