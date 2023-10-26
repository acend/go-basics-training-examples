package jokeapi

import "context"

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type JokeGetter interface {
	Get(ctx context.Context) (*Joke, error)
}

type JokeAdder interface {
	Add(ctx context.Context, joke *Joke) error
}

type JokeStore interface {
	JokeGetter
	JokeAdder
}
