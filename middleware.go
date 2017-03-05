package hyperdrive

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

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
