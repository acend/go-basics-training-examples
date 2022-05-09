package main

import (
	"fmt"
	"net/http"
	"os"

	api "user-api"
)

func main() {
	userStore := api.NewMemoryUserStore()

	userHandler := api.NewUserHandler(userStore)

	mux := http.NewServeMux()
	mux.Handle("/users/", userHandler)
	mux.Handle("/users", userHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
