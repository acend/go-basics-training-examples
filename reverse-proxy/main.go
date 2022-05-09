package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// the url where we want to redirect requests to
	target, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// recive requests on :8081, add a header X-Foo: bar and then forward the request to target
	err = http.ListenAndServe(":8081", addHeader("X-Auth", "xyz-token-xyz", forwardTo(target)))
	if err != nil {
		log.Fatal(err)
	}
}

func forwardTo(target *url.URL) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host
		proxy.ServeHTTP(w, r)
	})
}

func addHeader(header, value string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set(header, value)
		next.ServeHTTP(w, r)
	})
}
