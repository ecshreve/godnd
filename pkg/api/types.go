package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samsarahq/go/oops"
)

//go:generate genny

// ResourceList is a list of APIReferences for an endpoint.
type ResourceList struct {
	Count     int            `json:"count"`
	Resources []APIReference `json:"results"`
}

// APIReference is the basic data about an API resource.
type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (c *Client) GetAbilityScores(ctx context.Context) (*ResourceList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ability-scores", c.BaseURL), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request")
	}

	req = req.WithContext(ctx)

	res := ResourceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get ability-scores")
	}

	return &res, nil
}

func (c *Client) GetAbilityScoreByIndex(ctx context.Context, index string) (*apiAbilityScore, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ability-scores/%s", c.BaseURL, index), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request")
	}

	req = req.WithContext(ctx)

	res := apiAbilityScore{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get ability score")
	}

	return &res, nil
}
