package client

import (
	"context"
	"fmt"

	"github.com/samsarahq/go/oops"
)

type DamageType struct {
	Index       string
	Name        string
	Description []string
}

type apiDamageType struct {
	APIReference
	Description []string `json:"desc"`
}

func toDamageType(d apiDamageType) *DamageType {
	return &DamageType{
		Index:       d.Index,
		Name:        d.Name,
		Description: d.Description,
	}
}

func resourceURLFromIndex(index string) ResourceURL {
	return ResourceURL(fmt.Sprintf("/api/damage-types/%s", index))
}

func (c *Client) GetDamageType(ctx context.Context, index string) (*DamageType, error) {
	res := &apiDamageType{}

	if err := c.GetResource(ctx, resourceURLFromIndex(index), res); err != nil {
		return nil, oops.Wrapf(err, "unable to get resource for index: %s", index)
	}

	return toDamageType(*res), nil
}

func (c *Client) GetDamageTypes(ctx context.Context, indices []string) ([]*DamageType, error) {
	var damageTypes []*DamageType

	for _, index := range indices {
		damageType, err := c.GetDamageType(ctx, index)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to get damage types for indices: %v", indices)
		}

		damageTypes = append(damageTypes, damageType)
	}

	return damageTypes, nil
}

func (c *Client) GetAllDamageTypes(ctx context.Context) ([]*DamageType, error) {

	resList, err := c.GetResourceList(ctx, "damage-types")
	if err != nil {
		return nil, oops.Wrapf(err, "unable to get damage type list")
	}

	var damageTypes []*DamageType
	for _, apiReference := range resList.Resources {
		d := &apiDamageType{}
		if err := c.GetResourceForAPIReference(ctx, apiReference, d); err != nil {
			return nil, oops.Wrapf(err, "unable to get resource for APIReference: %v", apiReference)
		}

		damageTypes = append(damageTypes, toDamageType(*d))
	}

	return damageTypes, nil
}
