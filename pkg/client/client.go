package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samsarahq/go/oops"
)

// Client is a simple wrapper around an http.Client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a new Client with a timeout of one minute.
func NewClient() *Client {
	return &Client{
		BaseURL: getBaseURL(),
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
