package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	var addMode bool
	flag.IntVar(&delay, "delay", 2, "Specify a delay for the punchline")
	flag.StringVar(&url, "url", "https://official-joke-api.appspot.com/random_joke", "The URL that will be requested")
	flag.BoolVar(&addMode, "add", false, "Add a new joke")
	flag.Parse()

	if addMode {
		err := addJoke(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Exit the program, because we are in "add mode"
		os.Exit(0)
	}
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

func readJokeFromStdin() (joke, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Joke Setup: ")
	setupInput, err := reader.ReadString('\n')
	if err != nil {
		return joke{}, err
	}
	fmt.Print("Punchline: ")
	punchlineInput, err := reader.ReadString('\n')
	if err != nil {
		return joke{}, err
	}

	joke := joke{
		Setup:     strings.TrimSpace(setupInput),
		Punchline: strings.TrimSpace(punchlineInput),
	}
	return joke, nil
}

func addJoke(url string) error {
	joke, err := readJokeFromStdin()
	if err != nil {
		return err
	}

	data, err := json.Marshal(joke)
	if err != nil {
		return err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Unexpected status code. got: %v want: 201", res.StatusCode)
	}
	return nil
}
