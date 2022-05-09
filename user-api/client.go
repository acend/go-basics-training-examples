package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var _ UserStore = &Client{}

type Client struct {
	BaseURL string
	Client  *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "http://localhost:8080",
		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) Get(name string) (*User, error) {
	user := &User{}
	err := c.request("GET", "/users/"+name, nil, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) List() ([]User, error) {
	users := []User{}
	err := c.request("GET", "/users", nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) Delete(name string) error {
	err := c.request("DELETE", "/users/"+name, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Create(u *User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	fmt.Printf("post: '%s'\n", data)
	err = c.request("POST", "/users", bytes.NewBuffer(data), nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(method string, path string, input io.Reader, data any) (err error) {
	url := c.BaseURL + path

	req, err := http.NewRequest(method, url, input)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return fmt.Errorf("http request failed with code %d", resp.StatusCode)
	}

	if data == nil {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	return nil
}
