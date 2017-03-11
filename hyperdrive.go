package hyperdrive

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	slugify "github.com/metal3d/go-slugify"
)

// API is a logical collection of one or more endpoints, connecting requests
// to the response handlers using a gorlla mux Router.
type API struct {
	Name      string
	Desc      string
	Router    *mux.Router
	Server    *http.Server
	Root      *RootResource
	conf      Config
	endpoints []Endpoint
}

// NewAPI creates an instance of API, with an initialized Router, Config, Server, and RootResource.
func NewAPI(name string, desc string) API {
	api := API{
		Name:   name,
		Desc:   desc,
		Router: mux.NewRouter(),
		conf:   NewConfig(),
	}
	api.Root = NewRootResource(api)
	api.Router.Handle("/", api.DefaultMiddlewareChain(api.Root)).Methods("GET")
	api.Server = &http.Server{
		Handler:      api.Router,
		Addr:         api.conf.GetPort(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return api
}

// AddEndpoint registers endpoints, ensuring that endpoints automatically
// respond with a 405 error if the endpoint does not support a particular
// HTTP method.
func (api *API) AddEndpoint(e Endpointer) {
	api.Root.AddEndpointer(e)
	api.Router.Handle(e.GetPath(), api.DefaultMiddlewareChain(NewMethodHandler(e)))
}

// GetMediaType returns a media type string, sans any content-type extension (e.g. json),
// based on the name of the API, the Endpoint, and the Endpoint's version. The Media Type
// produced will be used for Content Negotiation, via the Accept header, as well as routing
// to the appropriate endpoint, when the media type appears in the request headers (e.g.
// Accept and Content-Type). It will also be used, after content negotation in the
// Content-Type response header.
func (api *API) GetMediaType(e Endpointer) string {
	return fmt.Sprintf("application/vnd.%s.%s.%s", slug(api.Name), slug(e.GetName()), e.GetVersion())
}

// Start starts the configured http server, listening on the configured Port
// (default: 5000). Set the PORT environment variable to change this.
func (api *API) Start() {
	log.Printf("Hyperdrive API starting on PORT %d in ENVIRONMENT %s", api.conf.Port, api.conf.Env)
	log.Fatal(api.Server.ListenAndServe())
}

func slug(s string) string {
	return strings.ToLower(slugify.Marshal(s))
}
