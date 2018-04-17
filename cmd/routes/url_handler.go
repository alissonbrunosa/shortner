package routes

import (
	"github.com/alissonbrunosa/shortner/internal/services"
	"net/http"
	"fmt"
)

type URLHandler struct {
	Service *services.URL
}

func (this *URLHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	link, ok := query["link"]
	if ok {
		link, err := this.Service.Create(link[0])

		if err != nil {
			writer.WriteHeader(500)
			fmt.Fprint(writer, "<h1>Ops...</h1><h3>Deu Ruim!</h3>")
			return
		}

		writer.WriteHeader(201)
		fmt.Fprintf(writer, "<h1><a href='%s' target='_blank' >Click!</a></h1>", link)
		return
	} else {
		writer.WriteHeader(422)
		fmt.Fprint(writer, "<h1>Ops...</h1><h3>Params 'link' is required</h3>")
		return
	}
}

func NewURLHandler() *URLHandler {
	return &URLHandler{
		Service: services.NewURLService(),
	}
}
