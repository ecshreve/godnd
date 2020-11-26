package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/samsarahq/go/oops"
)

// BaseURLV1 is the base url defined here: https://www.dnd5eapi.co/docs/#base.
const BaseURLV1 = "https://www.dnd5eapi.co/api/"

func getBaseURL() string {
	if os.Getenv("GODND_ENV") == "dev" {
		return "http://localhost:3000/api/"
	}

	return BaseURLV1
}

func getHTTPRequest(baseURL, resName string, resIndex *string) (*http.Request, error) {
	var reqURL string
	if resIndex != nil {
		reqURL = fmt.Sprintf("%s%s/%s", baseURL, resName, *resIndex)
	} else {
		reqURL = fmt.Sprintf("%s%s", baseURL, resName)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to build request for resource: %s and index: %s", resName, *resIndex)
	}

	return req, err
}

func allHelper(ctx context.Context, c *Client, resName string) (*ResourceList, error) {
	res := ResourceList{}
	req, err := getHTTPRequest(c.BaseURL, resName, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if len(res.Resources) == 0 {
		return nil, oops.Errorf("zero items returned for resource list: %s", resName)
	}

	return &res, nil
}

func byIndexHelper(ctx context.Context, c *Client, v interface{}, resName, resIndex string) error {
	req, err := getHTTPRequest(c.BaseURL, resName, &resIndex)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &v); err != nil {
		return err
	}

	return nil
}
