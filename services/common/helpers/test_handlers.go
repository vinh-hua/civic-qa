package helpers

import (
	"net/http"
	"net/http/httptest"
)

func makeReq(method, contentType string) *http.Request {
	req := httptest.NewRequest(method, "/path", nil)
	req.Header.Set("content-type", contentType)
	return req
}

// func TestHandlers(t *testing.T) {
// 	cases := []struct {
// 		r           *http.Request
// 		method      string
// 		contentType string

// 		expStatus int
// 		expErr    error
// 	}{
// 		{
// 			 r: makeReq("GET", "")
// 		}
// 	}
// }
