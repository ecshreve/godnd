package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samsarahq/go/oops"
)

// GetClasses returns the ResourceList for the /classes endpoint.
func (c *Client) GetClasses(ctx context.Context) (*ResourceList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/classes", c.BaseURL), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request")
	}

	req = req.WithContext(ctx)

	res := ResourceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get classes")
	}

	return &res, nil
}
