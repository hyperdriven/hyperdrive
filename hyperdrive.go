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
	return api
}

// AddEndpoint registers endpoints, ensuring that endpoints automatically
// respond with a 405 error if the endpoint does not support a particular
// HTTP method.
func (api *API) AddEndpoint(e Endpointer) {
	api.Root.AddEndpoint(e)
	api.Router.Handle(e.GetPath(), api.DefaultMiddlewareChain(NewMethodHandler(e))).
		Headers("Accept", api.GetContentTypeJSON(e)).
		Headers("Accet", api.GetContentTypeXML(e))
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

// GetContentTypeJSON returns the json Content-Type an endpoint can accept and
// respond with. The Content-Type will include the versioned vendor Media
// Type returned by API.GetMediaType() with a json extension.
func (api *API) GetContentTypeJSON(e Endpointer) string {
	return fmt.Sprintf("%s.json", api.GetMediaType(e))
}

// GetContentTypeXML returns the xml Content-Type an endpoint can accept and
// respond with. The Content-Type will include the versioned vendor Media
// Type returned by API.GetMediaType() with an xml extension.
func (api *API) GetContentTypeXML(e Endpointer) string {
	return fmt.Sprintf("%s.xml", api.GetMediaType(e))
}

// GetContentTypes returns a slice of Content-Types the endpoint can accept and
// respond with. The Content-Types will include both the versioned vendor Media
// Type returned by API.GetMediaType() for both json and xml.
func (api *API) GetContentTypes(e Endpointer) []string {
	return []string{api.GetContentTypeJSON(e), api.GetContentTypeXML(e)}
}

// Start starts the configured http server, listening on the configured Port
// (default: 5000). Set the PORT environment variable to change this.
func (api *API) Start() {
	log.Printf("Hyperdriven API: %s starting on: http://0.0.0.0:%d in: %s", api.Name, conf.Port, conf.Env)
	log.Fatal(api.Server.ListenAndServe())
}

func slug(s string) string {
	return strings.ToLower(slugify.Marshal(s))
}
