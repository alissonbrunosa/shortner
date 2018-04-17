package main

import (
	"github.com/alissonbrunosa/shortner/cmd/routes"
	"log"
	"net/http"
)

func initServer() *routes.Server {
	return &routes.Server{
		Dispatcher: routes.NewDispatcher(),
    URLHandler: routes.NewURLHandler(),
	}
}

func main() {
  server := initServer()
	log.Fatal(http.ListenAndServe(":8080", server.GetRouter()))
}
