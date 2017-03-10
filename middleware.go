package hyperdrive

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// DefaultMiddlewareChain wraps the given http.Handler in the following chain
// of middleware: LoggingMiddleware, RecoveryMiddleware.
func (api *API) DefaultMiddlewareChain(h http.Handler) http.Handler {
	return api.CompressionMiddleware(api.LoggingMiddleware(api.RecoveryMiddleware(h)))
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

// CompressionMiddleware wraps the given http.Handler and returns a gzipped response if
// the client requests it with the Accept-Encoding header. The compression level is set
// to to 1, by default. You can configure this though the
// GZIP_LEVEL environment variable, and set it to an integer between -2 and 9.
//
// Following zlib, levels range from 1 (Best Speed) to 9 (Best Compression); higher
// levels typically run slower but compress more.
//
// -1 is the Default Compression level, and is also used if an invalid value is
// configured via GZIP_LEVEL.
//
// 0 attemps no compression, and only adds the necessary DEFLATE framing.
//
// -2 disables Lempel-Ziv match searching and only performs Huffman entropy
// encoding. This is useful when compressing data that has already been compressed
// with an LZ style algorithm, such as Snappy or LZ4.
//
// More info can be found in the docs for the compress/flate package:
// https://golang.org/pkg/compress/flate/
func (api *API) CompressionMiddleware(h http.Handler) http.Handler {
	return handlers.CompressHandlerLevel(h, api.conf.GzipLevel)
}

// MethodOverrideMiddleware allows clients who can not perform native PUT, PATCH,
// or DELETE requests to specify the HTTP method in the X-HTTP-Method-Override
// header. The header name is case sensitive.
func (api *API) MethodOverrideMiddleware(h http.Handler) http.Handler {
	return handlers.HTTPMethodOverrideHandler(h)
}
