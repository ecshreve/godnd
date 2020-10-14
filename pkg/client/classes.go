package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samsarahq/go/oops"
)

// Class is the API representation of a DnD class.
type Class struct {
	APIReference
	HitDie             int            `json:"hit_die"`
	ProficiencyChoices []Choice       `json:"proficiency_choices"`
	Proficiencies      []APIReference `json:"proficiencies"`
	SavingThrows       []APIReference `json:"saving_throws"`
	StartingEquipment  string         `json:"starting_equipment"`
	ClassLevels        string         `json:"class_levels"`
	SubClasses         []APIReference `json:"subclasses"`
	SpellCasting       string         `json:"spellcasting"`
	Spells             string         `json:"spells"`
}

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

// GetClassByIndex returns the Class for the given index.
func (c *Client) GetClassByIndex(ctx context.Context, index string) (*Class, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/classes/%s", c.BaseURL, index), nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request")
	}

	req = req.WithContext(ctx)

	res := Class{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, oops.Wrapf(err, "unable to get class")
	}

	return &res, nil
}
