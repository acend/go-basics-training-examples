package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type joke struct {
	Type      string
	Setup     string
	Punchline string
}

func main() {
	var delay int
	var url string
	flag.IntVar(&delay, "delay", 2, "Specify a delay for the punchline")
	flag.StringVar(&url, "url", "https://official-joke-api.appspot.com/random_joke", "The URL that will be requested")
	flag.Parse()

	fmt.Println("Requesting joke from", url)
	joke, err := requestJoke(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(joke.Setup)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println(joke.Punchline)
}

func requestJoke(url string) (joke, error) {
	res, err := http.Get(url)
	if err != nil {
		return joke{}, fmt.Errorf("Error making http request: %w\n", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return joke{}, fmt.Errorf("Error reading response body: %w\n", err)
	}

	j := joke{}
	err = json.Unmarshal(body, &j)
	if err != nil {
		return joke{}, fmt.Errorf("Error decoding body: %w\n", err)
	}
	return j, nil
}
