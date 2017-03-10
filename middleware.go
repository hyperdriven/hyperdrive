package hyperdrive

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// DefaultMiddlewareChain wraps the given http.Handler in the following chain
// of middleware: LoggingMiddleware, RecoveryMiddleware.
func (api *API) DefaultMiddlewareChain(h http.Handler) http.Handler {
	return api.LoggingMiddleware(api.RecoveryMiddleware(h))
}

// LoggingMiddleware wraps the given http.Handler and outputs requests in Apache-style
// Combined Log format. All logging is done to STDOUT only.
func (api *API) LoggingMiddleware(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

// RecoveryMiddleware wraps the given http.Handler and recovers from panics. It wil log
// the stacktrace if HYPERDRIVE_ENVIRONMENT env var is not set to "production".
func (api *API) RecoveryMiddleware(h http.Handler) http.Handler {
	opt := handlers.PrintRecoveryStack(api.conf.Env != "production")
	return handlers.RecoveryHandler(opt)(h)
}
