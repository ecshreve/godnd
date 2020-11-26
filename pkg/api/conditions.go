package api

import (
	"context"

	"github.com/samsarahq/go/oops"
)

type ResourceInterface interface {
}

type Resource struct {
	EndpointName string
}

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

func (c *Client) ConditionAll(ctx context.Context) ([]*Condition, error) {
	res, err := allHelper(ctx, c, "conditions")
	if err != nil {
		return nil, oops.Wrapf(err, "unable to get resource list for conditions")
	}

	resList := []*Condition{}
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
	res := apiCondition{}
	if err := byIndexHelper(ctx, c, &res, "conditions", index); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource: %s with index: %s", "conditions", index)
	}

	return getConditionFromAPICondition(&res), nil
}
