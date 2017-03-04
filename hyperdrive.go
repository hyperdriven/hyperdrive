// Package hyperdrive is an opinonated micro-framework for writing hypermedia
// APIs. It attempts to embrace the best of Hypermedia, especially
// the seperation of client and server, encapsulated in the
// principle of HATEOAS (HTTP as the Engine of Application State).
//
// Hyperdrive APIs are resource-oriented, make heavy use of
// `http.Handler` middleware patterns, and takes advantage of
// HTTP verbs, headers, and other transport specific features, as
// much as possible. Other than that, it assumes nothing about how
// you store and retrieve your endpoint's hypermedia respresentations.
package hyperdrive

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// API is a logical collection of one or more endpoints, connecting requests
// to the response handlers using a gorlla mux Router.
type API struct {
	Router    *mux.Router
	endpoints []Endpoint
}

// NewAPI creates an instance of an API with an initialized Router.
func NewAPI() API {
	return API{Router: mux.NewRouter()}
}

// AddEndpoint registers endpoints, ensuring that endpoints automatically
// respond with a 405 error if the endpoint does not support a particular
// HTTP method.
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

// GetHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Get(). If this is not implemented,
// GET requests will be responded to with a `405 Method Not Allowed`
// error.
type GetHandler interface {
	Get(http.ResponseWriter, *http.Request)
}

// PostHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Post(). If this is not implemented,
// POST requests will be responded to with a `405 Method Not Allowed`
// error.
type PostHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

// PutHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Put(). If this is not implemented,
// PUT requests will be responded to with a `405 Method Not Allowed`
// error.
type PutHandler interface {
	Put(http.ResponseWriter, *http.Request)
}

// PatchHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Patch(). If this is not implemented,
// PATCH requests will be responded to with a `405 Method Not Allowed`
// error.
type PatchHandler interface {
	Patch(http.ResponseWriter, *http.Request)
}

// DeleteHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Delete(). If this is not implemented,
// DELETE requests will be responded to with a `405 Method Not Allowed`
// error.
type DeleteHandler interface {
	Delete(http.ResponseWriter, *http.Request)
}

// OptionsHandler interface is satisfied if the endpoint has implemented
// a http.Handler method called Options(). If this is not implemented,
// OPTIONS requests will be responded to with a `200 OK` and the `Allow`
// header will be set with a list of all the methods your endpoint does
// support.
type OptionsHandler interface {
	Options(http.ResponseWriter, *http.Request)
}

// Endpointer interface provides flexibility in how endpoints are created
// allowing for expressiveness how developers make use of this package.
type Endpointer interface {
	GetName() string
	GetDesc() string
	GetPath() string
}

// Endpoint is a basic implementation of the Endpointer interface and
// can be used directly if desired.
type Endpoint struct {
	EndpointName string
	EndpointDesc string
	EndpointPath string
}

// GetName satisfies part of the Endpointer interface and returns a
// string containing the name of the endpoint.
func (e *Endpoint) GetName() string {
	return e.EndpointName
}

// GetDesc satisfies part of the Endpointer interface and returns a
// string containing the description of the endpoint.
func (e *Endpoint) GetDesc() string {
	return e.EndpointDesc
}

// GetPath satisfies part of the Endpointer interface and returns a
// string containing the path of the endpoint, used by the Router.
//
// This string can contain named segmets, regex, and other features as
// described here: http://www.gorillatoolkit.org/pkg/mux
func (e *Endpoint) GetPath() string {
	return e.EndpointPath
}

// NewEndpoint creates an instance of Endpoint.
func NewEndpoint(name string, desc string, path string) *Endpoint {
	return &Endpoint{EndpointName: name, EndpointDesc: desc, EndpointPath: path}
}
