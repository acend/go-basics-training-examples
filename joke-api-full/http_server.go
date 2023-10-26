package jokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type JokeHandler struct {
	jokeStore JokeStore
}

func NewJokeHandler(jokeStore JokeStore) *JokeHandler {
	return &JokeHandler{
		jokeStore: jokeStore,
	}
}

func (j *JokeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		j.jokeGetHandler(w, r)
	case "POST":
		j.jokePostHandler(w, r)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (j *JokeHandler) jokeGetHandler(w http.ResponseWriter, r *http.Request) {
	joke, err := j.jokeStore.Get(r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(joke)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func (j *JokeHandler) jokePostHandler(w http.ResponseWriter, r *http.Request) {
	const MAX_JOKE_SIZE = 1_000_000

	body, err := io.ReadAll(io.LimitReader(r.Body, MAX_JOKE_SIZE))
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	joke := &Joke{}

	err = json.Unmarshal(body, joke)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if joke.Setup == "" {
		http.Error(w, "setup is empty", http.StatusBadRequest)
		return
	}

	if joke.Punchline == "" {
		http.Error(w, "punchline is empty", http.StatusBadRequest)
		return
	}

	err = j.jokeStore.Add(r.Context(), joke)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
