package hyperdrive

import (
	"fmt"
	"strings"
)

// GetMediaType returns a media type string, sans any content-type extension (e.g. json),
// based on the name of the API, the Endpoint, and the Endpoint's version. The Media Type
// produced will be used for Content Negotiation, via the Accept header, as well as routing
// to the appropriate endpoint, when the media type appears in the request headers (e.g.
// Accept and Content-Type). It will also be used, after content negotation in the
// Content-Type response header.
func GetMediaType(api API, e Endpointer) string {
	return fmt.Sprintf("application/vnd.%s.%s.%s", slug(api.Name), slug(e.GetName()), e.GetVersion())
}

// GetContentTypeJSON returns the json Content-Type an endpoint can accept and
// respond with. The Content-Type will include the versioned vendor Media
// Type returned by API.GetMediaType() with a json extension.
func GetContentTypeJSON(api API, e Endpointer) string {
	return fmt.Sprintf("%s.json", GetMediaType(api, e))
}

// GetContentTypeXML returns the xml Content-Type an endpoint can accept and
// respond with. The Content-Type will include the versioned vendor Media
// Type returned by API.GetMediaType() with an xml extension.
func GetContentTypeXML(api API, e Endpointer) string {
	return fmt.Sprintf("%s.xml", GetMediaType(api, e))
}

// GetContentTypes returns a slice of Content-Types that the endpoint can accept
// and respond with. The Content-Types will include both the versioned vendor
// Media Type returned by API.GetMediaType() for both json and xml.
func GetContentTypes(api API, e Endpointer) []string {
	return []string{GetContentTypeJSON(api, e), GetContentTypeXML(api, e)}
}

// GetContentTypesList returns a list of Content-Type strings that the endpoint can
// accept and respond with. The Content-Types will include both the versioned
// vendor Media Type returned by API.GetMediaType() for both json and xml.
func GetContentTypesList(api API, e Endpointer) string {
	return strings.Join(GetContentTypes(api, e), ",")
}
