package hyperdrive

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// LoggingMiddleware wraps http.Handlers and outputs requests in Apache-style
// Combined Log format. All logging is done to STDOUT only.
func LoggingMiddleware(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}
