package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

type joke struct {
	Type      string
	Setup     string
	Punchline string
}

var jokes []joke

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		n := rand.Intn(len(jokes))
		joke := jokes[n]
		data, err := json.Marshal(joke)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		j := joke{}
		err = json.Unmarshal(body, &j)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = addJoke(j)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	// Ensure we assign the jokes globally
	var err error
	jokes, err = readJokes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc("/", jokeHandler)

	fmt.Println("Serving API on http://localhost:8080")
	// ListenAndServe always returns a non-nil error.
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readJokes() ([]joke, error) {
	file, err := os.ReadFile("jokes.json")
	if err != nil {
		return nil, err
	}

	jokes := []joke{}
	if err := json.Unmarshal(file, &jokes); err != nil {
		return nil, err
	}
	return jokes, nil
}

func addJoke(j joke) error {
	jokes = append(jokes, j)
	data, err := json.Marshal(jokes)
	if err != nil {
		return err
	}
	err = os.WriteFile("jokes.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
