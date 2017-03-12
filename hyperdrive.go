package hyperdrive

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	slugify "github.com/metal3d/go-slugify"
)

var (
	hAPI API
)

// API is a logical collection of one or more endpoints, connecting requests
// to the response handlers using a gorlla mux Router.
type API struct {
	Name      string
	Desc      string
	Router    *mux.Router
	Server    *http.Server
	Root      *RootResource
	endpoints []Endpoint
}

// NewAPI creates an instance of API, with an initialized Router, Config, Server, and RootResource.
func NewAPI(name string, desc string) API {
	api := API{
		Name:   name,
		Desc:   desc,
		Router: mux.NewRouter(),
	}
	api.Root = NewRootResource(api)
	api.Router.Handle("/", api.DefaultMiddlewareChain(api.Root)).Methods("GET")
	api.Server = &http.Server{
		Handler:      api.Router,
		Addr:         conf.GetPort(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	hAPI = api
	return api
}

// AddEndpoint registers endpoints, ensuring that endpoints automatically
// respond with a 405 error if the endpoint does not support a particular
// HTTP method.
func (api *API) AddEndpoint(e Endpointer) {
	api.Root.AddEndpoint(e)
	api.Router.Handle(e.GetPath(), api.DefaultMiddlewareChain(NewMethodHandler(e))).
		Headers("Accept", GetContentTypeJSON(*api, e)).
		Headers("Accet", GetContentTypeXML(*api, e))
	log.Printf("Added hyperdriven Endpoint: %s http://0.0.0.0:%d%s", e.GetName(), conf.Port, e.GetPath())
	log.Printf("    Methods: %s", GetMethodsList(e))
	log.Printf("    Media Types: %s", GetContentTypesList(*api, e))

}

// Start starts the configured http server, listening on the configured Port
// (default: 5000). Set the PORT environment variable to change this.
func (api *API) Start() {
	log.Printf("Starting hyperdriven API (%s): %s http://0.0.0.0:%d", conf.Env, api.Name, conf.Port)
	log.Fatal(api.Server.ListenAndServe())
}

func slug(s string) string {
	return strings.ToLower(slugify.Marshal(s))
}
