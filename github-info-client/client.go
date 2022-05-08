package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	Name        string `json:"name"`
	Stars       int    `json:"stargazers_count"`
	Issues      int    `json:"open_issues"`
	Description string `json:"description"`
}

type User struct {
	Name      string `json:"login"`
	FullName  string `json:"name"`
	Followers int    `json:"followers"`
}

type Client struct {
	Debug   bool
	Token   string
	BaseURL string
	Client  *http.Client
}

func NewClient() *Client {
	return &Client{
		Debug:   false,
		BaseURL: "https://api.github.com",
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) GetUser(name string) (*User, error) {
	user := &User{}
	err := c.request("GET", "/users/"+name, nil, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetRepo(org string, name string) (*Repository, error) {
	repo := &Repository{}
	err := c.request("GET", "/repos/"+org+"/"+name, nil, repo)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (c *Client) request(method string, path string, input io.Reader, data any) (err error) {
	url := c.BaseURL + path

	fmt.Println(url)
	req, err := http.NewRequest(method, url, input)
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		return fmt.Errorf("http request failed with code %d", resp.StatusCode)
	}

	info, ok := resp.Header["X-Ratelimit-Remaining"]
	if c.Debug && ok {
		log.Printf("remaining requests: %s", info)
	}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	return nil
}
