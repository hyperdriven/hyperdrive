package hyperdrive

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

// ContentEncoder interface wraps the details of encoding response bodies to
// support automatic Content Negotiation.
type ContentEncoder interface {
	Encode(interface{}) error
}

// NullEncoder is an implementation of ContentEncoder, and is the default
// encoder used when Content Negotiation has falied. It produces a 406
// NOT ACCEPTABLE error when it's Encode() function is run.
type NullEncoder struct{}

// Encode returns a 406 NOT ACCEPTABLE error.
func (enc NullEncoder) Encode(v interface{}) error {
	return errors.New(http.StatusText(http.StatusNotAcceptable))
}

// JSONEncoder is an implementation of ContentEncoder and wraps the Encoder
// found in encoding/json package.
type JSONEncoder struct {
	Encoder *json.Encoder
}

// Encode encodes input as json text or returns an error.
func (enc JSONEncoder) Encode(v interface{}) error {
	return enc.Encoder.Encode(v)
}

// XMLEncoder is an implementation of ContentEncoder and wraps the Encoder
// found in encoding/xml package.
type XMLEncoder struct {
	Encoder *xml.Encoder
}

// Encode encodes input as xml text or returns an error.
func (enc XMLEncoder) Encode(v interface{}) error {
	return enc.Encoder.Encode(v)
}

// GetEncoder returns the correct ContentEncoder, determined by the Accept
// header, to support automatic Content Negotiation.
func GetEncoder(rw http.ResponseWriter, accept string) (ContentEncoder, http.ResponseWriter) {
	if strings.HasSuffix(accept, "json") {
		rw.Header().Set("Content-Type", accept)
		return JSONEncoder{json.NewEncoder(rw)}, rw
	}

	if strings.HasSuffix(accept, "xml") {
		rw.Header().Set("Content-Type", accept)
		return XMLEncoder{xml.NewEncoder(rw)}, rw
	}

	return NullEncoder{}, rw
}
