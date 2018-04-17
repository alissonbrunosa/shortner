package routes

import "net/http"

type Server struct {
	Dispatcher http.Handler
	URLHandler http.Handler
}

func (s *Server) GetRouter() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("/favicon.ico", func(w http.ResponseWriter, request *http.Request) {
		http.ServeFile(w, request, "resources/favicon.ico")
	})
	server.HandleFunc("/robots.txt", func(w http.ResponseWriter, request *http.Request) {
		http.ServeFile(w, request, "resources/robots.txt")
	})

	server.Handle("/create", s.URLHandler)
	server.Handle("/", s.Dispatcher)

	return server
}
