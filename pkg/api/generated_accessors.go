// Code generated by genny; DO NOT EDIT.

package api

import (
	"context"

	"github.com/samsarahq/go/oops"
)

// ConditionAll returns a slice of all Condition, or an error if unsuccessful.
func (c *Client) ConditionAll(ctx context.Context) ([]*Condition, error) {
	res, err := allHelper(ctx, c, "conditions")
	if err != nil {
		return nil, oops.Wrapf(err, "unable to get resource list for conditions")
	}

	resList := []*Condition{}
	for _, resource := range res.Resources {
		parsedRes, err := c.ConditionByIndex(ctx, resource.Index)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to get conditions with index %s for list", resource.Index)
		}
		resList = append(resList, parsedRes)
	}

	return resList, nil
}

// ConditionByIndex returns the Condition with the given index, or an error if unsuccessful.
func (c *Client) ConditionByIndex(ctx context.Context, index string) (*Condition, error) {
	res := apiCondition{}
	if err := byIndexHelper(ctx, c, &res, "conditions", index); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource: conditions with index: %s", index)
	}
	return res.convert(), nil
}

// AbilityScoreAll returns a slice of all AbilityScore, or an error if unsuccessful.
func (c *Client) AbilityScoreAll(ctx context.Context) ([]*AbilityScore, error) {
	res, err := allHelper(ctx, c, "ability-scores")
	if err != nil {
		return nil, oops.Wrapf(err, "unable to get resource list for ability-scores")
	}

	resList := []*AbilityScore{}
	for _, resource := range res.Resources {
		parsedRes, err := c.AbilityScoreByIndex(ctx, resource.Index)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to get ability-scores with index %s for list", resource.Index)
		}
		resList = append(resList, parsedRes)
	}

	return resList, nil
}

// AbilityScoreByIndex returns the AbilityScore with the given index, or an error if unsuccessful.
func (c *Client) AbilityScoreByIndex(ctx context.Context, index string) (*AbilityScore, error) {
	res := apiAbilityScore{}
	if err := byIndexHelper(ctx, c, &res, "ability-scores", index); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource: ability-scores with index: %s", index)
	}
	return res.convert(), nil
}

