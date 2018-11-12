package main

import (
	"os"

	"github.com/ZTao-z/gocloud-io/serverCode"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = PORT
	}

	flag.StringVarP(&port, "port", "p", PORT, "the server port")
	flag.Parse()

	server := service.NewServer()
	server.Run(":" + port)
}
