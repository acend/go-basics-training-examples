package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		fmt.Fprintf(w, "hello, your ip is %s\n", r.RemoteAddr)
	})

	manager := autocert.Manager{
		Cache:      autocert.DirCache("cert-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("go.dvob.ch"),
	}

	server := http.Server{
		Addr:      ":443",
		TLSConfig: manager.TLSConfig(),
	}

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}

}
