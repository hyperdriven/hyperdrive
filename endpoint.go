package hyperdrive

import "net/http"

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
