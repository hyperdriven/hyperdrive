package hyperdrive

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// QueryParams extracts the values from the request QueryString. It returns
// a url.Values object (essentially map[string][]string). If the
// request method is not GET, an empty url.Values is returned.
func QueryParams(r *http.Request) url.Values {
	if r.Method == "GET" {
		r.ParseForm()
		return r.Form
	}
	return url.Values{}
}

// BodyParams deserializes the input, and extracts the values from the request
// body. It returns a url.Values object (essentially map[string][]string). If
// the request method is GET, an empty url.Values is returned.
func BodyParams(r *http.Request) url.Values {
	if r.Method != "GET" {
		return url.Values{}
	}
	return url.Values{}
}

// PathParams extracts the values from the request path which match named
// params in the route. They are returned as url.Values for consistincey
// with http.Request.Form's behaviour.
func PathParams(r *http.Request) url.Values {
	var params = url.Values{}
	for k, v := range mux.Vars(r) {
		params.Add(k, v)
	}
	return params
}

// Params extracts the param values from all sources: query, body, and path -- in
// that order. Each subsequent source will overwrite values with the same key, to
// ensure API client intent is maintained in a consistent way.
func Params(r *http.Request) url.Values {
	var params = QueryParams(r)

	for k, values := range BodyParams(r) {
		for _, v := range values {
			params.Set(k, v)
		}
	}

	for k, v := range mux.Vars(r) {
		params.Set(k, v)
	}

	return params
}
