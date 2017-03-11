package hyperdrive

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/gorilla/handlers"
)

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
// allowing for expressiveness in how developers make use of the hyperdrive
// package.
type Endpointer interface {
	GetName() string
	GetDesc() string
	GetPath() string
	GetVersion() string
}

// Endpoint is a basic implementation of the Endpointer interface and
// can be used directly if desired.
type Endpoint struct {
	EndpointName    string
	EndpointDesc    string
	EndpointPath    string
	EndpointVersion *semver.Version
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

// GetVersion returns a string representing the version.
func (e *Endpoint) GetVersion() string {
	var v = fmt.Sprintf("v%d", e.EndpointVersion.Major())

	if (e.EndpointVersion.Major() >= 0 && e.EndpointVersion.Minor() != 0) || (e.EndpointVersion.Major() >= 0 && e.EndpointVersion.Patch() > 0) {
		v = fmt.Sprintf("%s%s%d", v, ".", e.EndpointVersion.Minor())
	}

	if e.EndpointVersion.Patch() != 0 {
		v = fmt.Sprintf("%s%s%d", v, ".", e.EndpointVersion.Patch())
	}

	if e.EndpointVersion.Prerelease() != "" {
		v = fmt.Sprintf("%s%s%s", v, "-", e.EndpointVersion.Prerelease())
	}

	return v
}

// NewEndpoint creates an instance of Endpoint.
func NewEndpoint(name string, desc string, path string, version string) *Endpoint {
	var (
		v   *semver.Version
		err error
	)
	v, err = semver.NewVersion(version)
	if err != nil {
		v, _ = semver.NewVersion("1")
	}
	return &Endpoint{EndpointName: name, EndpointDesc: desc, EndpointPath: path, EndpointVersion: v}
}

// GetMethods returns a slice of the methods an Endpoint supports.
func GetMethods(e Endpointer) []string {
	var methods = []string{"OPTIONS"}

	if _, ok := interface{}(e).(GetHandler); ok {
		methods = append(methods, "GET")
	}

	if _, ok := interface{}(e).(PostHandler); ok {
		methods = append(methods, "POST")
	}

	if _, ok := interface{}(e).(PutHandler); ok {
		methods = append(methods, "PUT")
	}

	if _, ok := interface{}(e).(PatchHandler); ok {
		methods = append(methods, "PATCH")
	}

	if _, ok := interface{}(e).(DeleteHandler); ok {
		methods = append(methods, "DELETE")
	}

	return methods
}

// GetMethodsList returns a list of the methods an Endpoint supports.
func GetMethodsList(e Endpointer) string {
	return strings.Join(GetMethods(e), ", ")
}

// NewMethodHandler sets the correct http.Handler for each method, depending on
// the interfaces the Enpointer supports. It returns an http.HandlerFunc, ready
// to be served directly, wrapped in other middleware, etc.
func NewMethodHandler(e Endpointer) http.HandlerFunc {
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
	return http.HandlerFunc(handler.ServeHTTP)
}
