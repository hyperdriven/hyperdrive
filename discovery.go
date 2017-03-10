package hyperdrive

import (
	"encoding/json"
	"net/http"
)

// Representation is a data structure representing the response output. The
// representation is used when automatically encoding responses based on the
// Content Type determined by content negotation.
type Representation map[string]interface{}

// RootResource contains information about the API and its Endpoints, and is
// the hypermedia respresentation returned by the Discovery URL endpoint for
// API clients to learn about the API.
type RootResource struct {
	Name      string
	Endpoints []Endpointer
}

// NewRootResource creates an instance of RootResource, based on the given API.
func NewRootResource(api API) *RootResource {
	return &RootResource{Name: api.Name}
}

// AddEndpointer adds Endpointers to the slice of Endpointers on an instance of RootResource.
func (root *RootResource) AddEndpointer(e Endpointer) {
	root.Endpoints = append(root.Endpoints, e)
}

// Present returns an Representation of the RootResource to describe the API
// for the Discovery URL.
func (root *RootResource) Present() Representation {
	return Representation{
		"resource":  "api",
		"name":      root.Name,
		"endpoints": root.endpointRepresentations(),
	}
}

func (root *RootResource) endpointRepresentations() []Representation {
	var endpoints = []Representation{}
	for _, e := range root.Endpoints {
		endpoints = append(endpoints, PresentEndpoint(e))
	}
	return endpoints
}

// PresentEndpoint returns a Representation to describe an Endpoint for the Discovery URL.
func PresentEndpoint(e Endpointer) Representation {
	return Representation{
		"name":    e.GetName(),
		"desc":    e.GetDesc(),
		"path":    e.GetPath(),
		"methods": GetMethods(e),
	}
}

// ServeHTTP satisfies the http.Handler interface and returns the hypermedia
// representation of the Discovery URL.
func (root *RootResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(root.Present())
}
