package hyperdrive

import (
	"net/http"

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
	api.Router.HandleFunc(e.GetPath(), NoMethodHandler(e))
}

type Endpoint struct {
	Name string
	Desc string
	Path string
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

type GetHandler interface {
	Get(http.ResponseWriter, *http.Request) http.HandlerFunc
}

func NoMethodHandler(endpoint Endpointer) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if h, ok := interface{}(endpoint).(GetHandler); ok {
				return h.Get(rw, r)
			} else {
				http.Error(rw, http.StatusText(405), 405)
			}
		}

	}
	return fn
}
