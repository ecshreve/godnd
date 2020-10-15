package client

import (
	"os"
)

// BaseURLV1 is the base url defined here: https://www.dnd5eapi.co/docs/#base.
const BaseURLV1 = "https://www.dnd5eapi.co/api/"

func getBaseURL() string {
	if os.Getenv("GODND_ENV") == "dev" {
		return "http://localhost:3000/api/"
	}

	return BaseURLV1
}

func (c *Client) getRawBaseURL() string {
	return c.BaseURL[:len(c.BaseURL)-5]
}
