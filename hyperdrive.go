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

type GetHandler interface {
	Get(http.ResponseWriter, *http.Request) http.HandlerFunc
}

type Endpointer interface {
	GetPath() string
	GetHandler
}

func (e *Endpoint) GetPath() string {
	return e.Path
}

func NewEndpoint(name string, desc string, path string) *Endpoint {
	return &Endpoint{Name: name, Desc: desc, Path: path}
}

func NoMethodHandler(endpoint Endpointer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if _, ok := interface{}(endpoint).(GetHandler); ok {
				endpoint.Get(rw, r)
				return
			} else {
				http.Error(rw, http.StatusText(405), 405)
				return
			}
		}

	}
}
