package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samsarahq/go/oops"
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

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return oops.Wrapf(err, "unable to make request")
	}
	defer res.Body.Close()

	// TODO: better status code handling.
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return oops.Errorf("response error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return oops.Wrapf(err, "unable to decode response bodybody")
	}

	return nil
}
