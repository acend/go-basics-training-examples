package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

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
		apiURL         = "https://official-joke-api.appspot.com/random_joke"
		jokeFile       = "jokes.json"
		punchlineDelay = time.Second * 5
		jokeSource     = "api"

		addJoke   bool
		setup     string
		punchline string
	)

	flag.StringVar(&apiURL, "url", apiURL, "api url which is used as joke store. only used if source is 'api'")
	flag.StringVar(&jokeFile, "file", jokeFile, "file which is used as joke store. only used if source is 'file'")
	flag.StringVar(&jokeSource, "src", jokeSource, "joke source to use (api, file)")

	flag.DurationVar(&punchlineDelay, "delay", punchlineDelay, "delay before printing the punchline")

	flag.BoolVar(&addJoke, "add", addJoke, "add a joke")
	flag.StringVar(&setup, "setup", setup, "The setup part of the joke when using the add joke option.")
	flag.StringVar(&punchline, "punchline", punchline, "The punchline part of the joke when using the add joke option.")

	flag.Parse()

	ctx := context.Background()

	var (
		jokeStore jokeapi.JokeStore
		err       error
	)

	// setup the joke store
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

	// add a joke if flag -add is set
	if addJoke {
		if setup != "" && punchline != "" {
			joke := &jokeapi.Joke{
				Setup:     setup,
				Punchline: punchline,
			}
			return jokeStore.Add(ctx, joke)
		}

		joke, err := readJokeFromStdin()
		if err != nil {
			return err
		}

		return jokeStore.Add(ctx, joke)
	}

	// print a joke
	joke, err := jokeStore.Get(ctx)
	if err != nil {
		return err
	}

	fmt.Println(joke.Setup)
	time.Sleep(punchlineDelay)
	fmt.Println(joke.Punchline)
	return nil
}

func readJokeFromStdin() (*jokeapi.Joke, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Joke Setup: ")
	setupInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fmt.Print("Punchline: ")
	punchlineInput, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	joke := &jokeapi.Joke{
		Setup:     strings.TrimSpace(setupInput),
		Punchline: strings.TrimSpace(punchlineInput),
	}
	return joke, nil
}
