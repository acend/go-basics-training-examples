package jokeapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apiURL string
	client *http.Client
}

var _ JokeStore = (*Client)(nil)

func NewClient(apiURL string) (*Client, error) {
	_, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiURL: apiURL,
		client: &http.Client{
			Timeout: time.Second * 20,
		},
	}, nil
}

// Get obtains a random joke from the API.
func (c *Client) Get(ctx context.Context) (*Joke, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api %s returned error code %d", c.apiURL, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	joke := &Joke{}
	err = json.Unmarshal(data, joke)
	if err != nil {
		return nil, err
	}

	return joke, nil
}

// Add adds a joke.
func (c *Client) Add(ctx context.Context, joke *Joke) error {
	data, err := json.Marshal(joke)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.apiURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("api %s returned error code %d", c.apiURL, resp.StatusCode)
	}

	return nil
}
