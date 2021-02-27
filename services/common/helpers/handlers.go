package helpers

import (
	"errors"
	"net/http"
)

func ExpectMethodAndContentType(r *http.Request, method string, contentType string) (status int, err error) {
	if r.Method != method {
		return http.StatusMethodNotAllowed, errors.New("Method Not Allowed")
	}
	if r.Header.Get("content-type") != contentType {
		return http.StatusUnsupportedMediaType, errors.New("Unsupported content-type")
	}

	return -1, nil
}
