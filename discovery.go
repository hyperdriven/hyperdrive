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
	Endpoints []EndpointResource `json:"endpoints" xml:"endpoints>endpoint"`
}

// EndpointResource contains information about and Endpoint, and is
// the hypermedia respresentation returned by the Discovery URL endpoint for
// API clients to learn about the Endpoint.
type EndpointResource struct {
	XMLName        xml.Name                `json:"-" xml:"endpoint"`
	Resource       string                  `json:"resource" xml:"-"`
	Name           string                  `json:"name" xml:"name,attr"`
	Path           string                  `json:"path" xml:"path,attr"`
	MethodsList    string                  `json:"-" xml:"methods,attr"`
	Methods        []string                `json:"methods" xml:"-"`
	MediaTypesList string                  `json:"-" xml:"media-types,attr"`
	MediaTypes     []string                `json:"media-types" xml:"-"`
	Desc           string                  `json:"description" xml:"description"`
	Params         []EndpointResourceParam `json:"params" xml:"params>param"`
}

// EndpointResourceParam contains information about endpoint parameters, and is
// part of the hypermedia representation returned by the Discovery URL endpoint
// for API clients to learn about input allowed (and/or required) by the
// Endpoint.
type EndpointResourceParam struct {
	XMLName      xml.Name `json:"-" xml:"param"`
	Name         string   `json:"name" xml:"name,attr"`
	Desc         string   `json:"description" xml:"description"`
	AllowedList  string   `json:"-" xml:"allowed,attr"`
	Allowed      []string `json:"allowed" xml:"-"`
	RequiredList string   `json:"-" xml:"required,attr"`
	Required     []string `json:"required" xml:"-"`
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
		Params:         createEndpointResourceParams(e),
	}
}

// NewEndpointResourceParam creates an instance of EndpointResourceParam from the given parsedParam.
func NewEndpointResourceParam(p parsedParam) EndpointResourceParam {
	return EndpointResourceParam{
		Name:         p.Name,
		Desc:         p.Desc,
		Allowed:      p.Allowed,
		AllowedList:  p.AllowedList(),
		Required:     p.Required,
		RequiredList: p.RequiredList(),
	}
}

func createEndpointResourceParams(e Endpointer) []EndpointResourceParam {
	var params = []EndpointResourceParam{}
	pp := parseEndpoint(e)
	for _, p := range pp {
		params = append(params, NewEndpointResourceParam(p))
	}
	return params
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
