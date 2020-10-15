package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samsarahq/go/oops"
)

type Object interface{}

// ResourceList is a list of APIReferences for an endpoint.
type ResourceList struct {
	Count     int            `json:"count"`
	Resources []APIReference `json:"results"`
}

// ResourceURL is a type alias for a string for readability purposes.
type ResourceURL string

// APIReference is the basic data about an API resource.
type APIReference struct {
	Index string      `json:"index"`
	Name  string      `json:"name"`
	URL   ResourceURL `json:"url"`
}

// Choice represents a selection from a list of resources.
type Choice struct {
	Choose int            `json:"choose"`
	Type   string         `json:"type"`
	From   []APIReference `json:"from"`
}

// GetResourceList returns the ResourceList for the given endpoint.
func (c *Client) GetResourceList(ctx context.Context, endpoint string) (*ResourceList, error) {
	resURL := ResourceURL(fmt.Sprintf("/api/%s", endpoint))

	res := &ResourceList{}
	if err := c.GetResource(ctx, resURL, res); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource list for endpoint: %s", endpoint)
	}

	return res, nil
}

// GetResource makes a request to the given resURL and unmarshals the response
// into the given object.
func (c *Client) GetResource(ctx context.Context, resURL ResourceURL, o Object) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.getRawBaseURL(), resURL), nil)
	if err != nil {
		return oops.Wrapf(err, "unable to build request for resource url: %s", resURL)
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &o); err != nil {
		return oops.Wrapf(err, "unable to get resource for request: %+v", req)
	}

	return nil
}

// GetResourceForAPIReference is a convenience wrapper around GetResource.
func (c *Client) GetResourceForAPIReference(ctx context.Context, ref APIReference, o Object) error {
	if err := c.GetResource(ctx, ref.URL, o); err != nil {
		return oops.Wrapf(err, "unable to get resource for api reference: %v", ref)
	}

	return nil
}
