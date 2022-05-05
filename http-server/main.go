package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	}
}

func loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("client=%s method=%s path=%s duration=%s",
			r.RemoteAddr, r.Method, r.URL.Path, time.Now().Sub(start))
	}
}

func run() error {
	handler := loggingMiddleware(pingHandler())
	return http.ListenAndServe(":8080", handler)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
