package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	jokeapi "joke-api"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		apiURL     = "https://official-joke-api.appspot.com/random_joke"
		jokeFile   = "jokes.json"
		jokeSource = "api"

		listenAddr = ":8080"
	)

	flag.StringVar(&apiURL, "url", apiURL, "api url which is used as joke store. only used if source is 'api'")
	flag.StringVar(&jokeFile, "file", jokeFile, "file which is used as joke store. only used if source is 'file'")
	flag.StringVar(&jokeSource, "src", jokeSource, "joke source to use (api, file)")

	flag.StringVar(&listenAddr, "addr", listenAddr, "listen address of the API server")

	flag.Parse()

	var (
		jokeStore jokeapi.JokeStore
		err       error
	)

	switch jokeSource {
	case "api":
		jokeStore, err = jokeapi.NewClient(apiURL)
		if err != nil {
			return err
		}
	case "file":
		jokeStore = jokeapi.NewFileStore(jokeFile)
	default:
		return fmt.Errorf("unknown joke source '%s'", jokeSource)
	}

	mux := http.NewServeMux()

	mux.Handle("/joke", jokeapi.NewJokeHandler(jokeStore))

	return http.ListenAndServe(listenAddr, mux)
}
