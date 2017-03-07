package hyperdrive

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// API is a logical collection of one or more endpoints, connecting requests
// to the response handlers using a gorlla mux Router.
type API struct {
	Router    *mux.Router
	Server    *http.Server
	conf      Config
	endpoints []Endpoint
}

// NewAPI creates an instance of an API with an initialized Router.
func NewAPI() API {
	api := API{Router: mux.NewRouter(), conf: NewConfig()}
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

	middleware := api.LoggingMiddleware(api.RecoveryMiddleware(http.HandlerFunc(handler.ServeHTTP)))
	api.Router.Handle(e.GetPath(), middleware)
}

// Start starts the configured http server, listening on the configured Port
// (default: 5000). Set the PORT environment variable to change this.
func (api *API) Start() {
	log.Printf("Hyperdrive API starting on PORT %d in ENVIRONMENT %s", api.conf.Port, api.conf.Env)
	log.Fatal(api.Server.ListenAndServe())
}
