package routes

import (
	"fmt"
	"github.com/alissonbrunosa/shortner/internal/services"
	"net/http"
)

type Dispatcher struct {
	Service *services.PathExtractor
	URLService *services.URL
}

func (this *Dispatcher) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(services.Hex(4))
	path := this.Service.Call(request.URL.Path)
	link, err := this.URLService.Find(path)
	if err != nil {
		writer.WriteHeader(500)
		fmt.Fprint(writer, "<h1>Ops...</h1><h3>Deu ruim!</h3>")
		return
	} 

  if link != "" {
		http.Redirect(writer, request, link, 301)
		return
	} 

	writer.WriteHeader(404)
	fmt.Fprint(writer, "<h1>Ops...</h1><h3>Page not found</h3>")
	return
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Service: services.NewPathExtractor(),
		URLService: services.NewURLService(),
	}
}
