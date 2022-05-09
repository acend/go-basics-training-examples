package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type User struct {
	Name      string `json:"login"`
	FullName  string `json:"name"`
	Followers int    `json:"followers"`
}

type Client struct {
	Debug  bool
	Token  string
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Debug: false,
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) GetUser(name string) (*User, error) {
	url := "https://api.github.com/users/" + name

	req, err := http.NewRequest("GET", url, nil)
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("http request failed with code %d", resp.StatusCode)
	}

	info, ok := resp.Header["X-Ratelimit-Remaining"]
	if c.Debug && ok {
		fmt.Println("remaining requests:", info[0])
	}

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	return user, nil
}

func main() {
	var (
		debug = false
	)

	flag.BoolVar(&debug, "debug", debug, "show additional information about rate limiting")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing argument: username")
		os.Exit(1)
	}

	userName := flag.Arg(0)

	githubClient := NewClient()
	githubClient.Debug = debug

	if token, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		githubClient.Token = token
	}

	user, err := githubClient.GetUser(userName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get user:", err)
		os.Exit(1)
	}

	fmt.Printf("name: %s\n", user.FullName)
	fmt.Printf("followers: %d\n", user.Followers)
}
