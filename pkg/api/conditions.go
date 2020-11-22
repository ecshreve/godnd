package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samsarahq/go/oops"
)

type Condition struct {
	Name        string
	Description []string
}

func getConditionFromAPICondition(a *apiCondition) *Condition {
	if a == nil {
		return nil
	}

	return &Condition{
		Name:        a.Name,
		Description: a.Desc,
	}
}

func (c *Client) ConditionsAll(ctx context.Context) ([]*Condition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sconditions", c.BaseURL), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request for resource list: %s", "conditions")
	}
	req = req.WithContext(ctx)

	res := ResourceList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource list: %s", "conditions")
	}

	if len(res.Resources) == 0 {
		return nil, oops.Errorf("zero items returned for resource list: %s", "conditions")
	}

	var resList []*Condition
	for _, resource := range res.Resources {
		parsedRes, err := c.ConditionByIndex(ctx, resource.Index)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to get resource: %s for list", resource.Index)
		}
		resList = append(resList, parsedRes)
	}

	return resList, nil
}

func (c *Client) ConditionByIndex(ctx context.Context, index string) (*Condition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sconditions/%s", c.BaseURL, index), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request for resource: %s and index: %s", "conditions", index)
	}
	req = req.WithContext(ctx)

	res := apiCondition{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource: %s with index: %s", "conditions", index)
	}

	parsedRes := getConditionFromAPICondition(&res)
	return parsedRes, nil
}
