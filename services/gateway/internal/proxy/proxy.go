package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func MustParse(rawurl string) *url.URL {
	url, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}

	return url
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Path = target.Path
	}

	return &httputil.ReverseProxy{Director: director}
}
