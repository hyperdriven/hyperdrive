package hyperdrive

import "net/http"

// GetErrorText helps ensure implementation details are not leaked in production
// environments. If this is production, it returns the http.StatusText for the
// given status code. If this is not production, the error message is returned
// to aid in debugging.
func GetErrorText(status int, err error) string {
	if conf.Env != "production" {
		return err.Error()
	}
	return http.StatusText(status)
}
