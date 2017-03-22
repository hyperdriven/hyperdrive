package hyperdrive

import (
	"reflect"
	"strings"

	"github.com/xtgo/set"
)

const (
	tagName = "param"
)

type parsedParam struct {
	Field    string
	Type     string
	Key      string
	Allowed  []string
	Required []string
}

func (p parsedParam) IsAllowed(method string) bool {
	return contains(p.Allowed, method)
}

func (p parsedParam) IsRequired(method string) bool {
	return contains(p.Required, method)
}

type parsedParams map[string]parsedParam

func (pp parsedParams) Allowed(method string) []string {
	var allowed []string
	for _, p := range pp {
		if p.IsAllowed(method) {
			allowed = append(allowed, p.Key)
		}
	}
	return allowed
}

func (pp parsedParams) Required(method string) []string {
	var required []string
	for _, p := range pp {
		if p.IsRequired(method) {
			required = append(required, p.Key)
		}
	}
	return required
}

func contains(haystack []string, needle string) bool {
	for _, m := range haystack {
		if m == needle {
			return true
		}
	}
	return false
}

func parseEndpoint(e Endpointer) parsedParams {
	var params = parsedParams{}
	t := reflect.TypeOf(e).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if _, ok := field.Tag.Lookup(tagName); ok {
			parsed := parseField(field)
			params[parsed.Key] = parsed
		}
	}
	return params
}

func parseField(field reflect.StructField) parsedParam {
	var (
		key      string
		tags     []string
		allowed  = []string{"GET", "POST", "PUT", "PATCH"}
		required = []string{}
	)
	t := field.Tag.Get(tagName)
	tags = strings.Split(t, ";")
	key, tags = tags[0], tags[1:]
	if key == "" {
		key = field.Name
	}

	for _, tag := range tags {
		pairs := strings.Split(tag, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "a":
			if v != "" {
				allowed = strings.Split(v, ",")
			}
		case "r":
			if v != "" {
				required = strings.Split(v, ",")
				allowed = set.Strings(allowed)
			}
		}
	}
	required = set.Strings(required)
	allowed = append(allowed, required...)
	allowed = set.Strings(allowed)
	return parsedParam{field.Name, field.Type.Name(), key, allowed, required}
}
