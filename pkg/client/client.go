package client

import (
	"net/http"
	"time"
)

// BaseURLV1 is the base url defined here: https://www.dnd5eapi.co/docs/#base.
const BaseURLV1 = "https://www.dnd5eapi.co/api/"

// Client is a simple wrapper around an http.Client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a new Client with a timeout of one minute.
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
