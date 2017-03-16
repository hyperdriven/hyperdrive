// Package hyperdrive is an opinonated micro-framework for writing hypermedia
// APIs. It attempts to embrace the best of Hypermedia, especially
// the separation of client and server, encapsulated in the
// principle of HATEOAS (HTTP as the Engine of Application State).
//
// Hyperdrive APIs are resource-oriented, make heavy use of
// `http.Handler` middleware patterns, and takes advantage of
// HTTP verbs, headers, and other transport specific features, as
// much as possible. Other than that, it assumes nothing about how
// you store and retrieve your endpoint's hypermedia respresentations.
package hyperdrive
