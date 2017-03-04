package hyperdrive

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	Router    *mux.Router
	endpoints []Endpoint
}

func NewAPI() API {
	return API{Router: mux.NewRouter()}
}

func (api *API) AddEndpoint(e Endpointer) {
	handler := make(handlers.MethodHandler)
	if h, ok := interface{}(e).(GetHandler); ok {
		handler["GET"] = http.HandlerFunc(h.Get)
	}

	if h, ok := interface{}(e).(PostHandler); ok {
		handler["POST"] = http.HandlerFunc(h.Post)
	}

	if h, ok := interface{}(e).(PutHandler); ok {
		handler["PUT"] = http.HandlerFunc(h.Put)
	}

	if h, ok := interface{}(e).(PatchHandler); ok {
		handler["PATCH"] = http.HandlerFunc(h.Patch)
	}

	if h, ok := interface{}(e).(DeleteHandler); ok {
		handler["DELETE"] = http.HandlerFunc(h.Delete)
	}

	if h, ok := interface{}(e).(OptionsHandler); ok {
		handler["OPTIONS"] = http.HandlerFunc(h.Options)
	}

	api.Router.HandleFunc(e.GetPath(), handler.ServeHTTP)
}

type GetHandler interface {
	Get(http.ResponseWriter, *http.Request)
}

type PostHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type PutHandler interface {
	Put(http.ResponseWriter, *http.Request)
}

type PatchHandler interface {
	Patch(http.ResponseWriter, *http.Request)
}

type DeleteHandler interface {
	Delete(http.ResponseWriter, *http.Request)
}

type OptionsHandler interface {
	Options(http.ResponseWriter, *http.Request)
}

type Endpoint struct {
	Name string
	Desc string
	Path string
	//GetHandler
	//PostHandler
	//PutHandler
	//PatchHandler
	//DeleteHandler
	//OptionsHandler
}

type Endpointer interface {
	GetPath() string
}

func (e *Endpoint) GetPath() string {
	return e.Path
}

func NewEndpoint(name string, desc string, path string) *Endpoint {
	return &Endpoint{Name: name, Desc: desc, Path: path}
}
