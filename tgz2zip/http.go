package main

import (
	"log"
	"net/http"
)

func convertServer() error {
	return http.ListenAndServe(":8080", http.HandlerFunc(convertHandler))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := tgzToZIP(r.Body, w)
	if err != nil {
		log.Println(err)
	}
}
