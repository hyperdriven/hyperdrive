package hyperdrive

import (
	"encoding/xml"
	"net/http"
)

// RootResource contains information about the API and its Endpoints, and is
// the hypermedia respresentation returned by the Discovery URL endpoint for
// API clients to learn about the API.
type RootResource struct {
	XMLName   xml.Name           `json:"-" xml:"api"`
	Resource  string             `json:"resource" xml:"-"`
	Name      string             `json:"name" xml:"name,attr"`
	Endpoints []EndpointResource `json:"endpoints" xml:"endpoints"`
}

// EndpointResource contains information about and Endpoint, and is
// the hypermedia respresentation returned by the Discovery URL endpoint for
// API clients to learn about the Endpoint.
type EndpointResource struct {
	XMLName        xml.Name `json:"-" xml:"endpoint"`
	Resource       string   `json:"resource" xml:"-"`
	Name           string   `json:"name" xml:"name,attr"`
	Path           string   `json:"path" xml:"path,attr"`
	MethodsList    string   `json:"-" xml:"methods,attr"`
	Methods        []string `json:"methods" xml:"-"`
	MediaTypesList string   `json:"-" xml:"media-types,attr"`
	MediaTypes     []string `json:"media-types" xml:"-"`
	Desc           string   `json:"description" xml:"description"`
}

// NewRootResource creates an instance of RootResource from the given API.
func NewRootResource(api API) *RootResource {
	return &RootResource{Resource: "api", Name: api.Name}
}

// NewEndpointResource creates an instance of EndpointResource from the given Endpointer.
func NewEndpointResource(e Endpointer) EndpointResource {
	return EndpointResource{
		Resource:       "endpoint",
		Name:           e.GetName(),
		Path:           e.GetPath(),
		MethodsList:    GetMethodsList(e),
		Methods:        GetMethods(e),
		MediaTypesList: GetContentTypesList(hAPI, e),
		MediaTypes:     GetContentTypes(hAPI, e),
		Desc:           e.GetDesc(),
	}
}

// AddEndpoint adds EndpointResources to the slice of Endpoints on an instance of RootResource.
func (root *RootResource) AddEndpoint(e Endpointer) {
	root.Endpoints = append(root.Endpoints, NewEndpointResource(e))
}

// ServeHTTP satisfies the http.Handler interface and returns the hypermedia
// representation of the Discovery URL.
func (root *RootResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	Respond(rw, r, 200, root)
}
